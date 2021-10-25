/*
	Main file for the logging module
    Copyright (C) 2021 Jack C <jack@chaker.net>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/json"
	"os"

	"github.com/chakernet/bizkit/common/amqp"
	"github.com/chakernet/bizkit/common/handler"
	"github.com/chakernet/bizkit/common/redis"
	"github.com/chakernet/bizkit/common/util"
	"github.com/chakernet/bizkit/logging/events"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/joho/godotenv"
	_amqp "github.com/streadway/amqp"
)

var (
	log = util.Logger{
		Name: "main",
	}
)

func main() {
	log.Info("Initializing...")
	// Load Env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Failed to load env: %s", err)
	}
	token := os.Getenv("BOT_TOKEN")

	// Connect to redis
	rdb := redis.Redis{}
	redconn := rdb.Connect()
	defer redconn.Close()
	log.Info("Initialized Redis Connection")

	// Connect to amqp
	amconn, err := amqp.Connect()
	if err != nil {
		log.Fatal("Error connecting to rabbitmq: %s", err)
	}
	defer amconn.Close()
	ch, err := amqp.Channel(amconn)
	if err != nil {
		log.Fatal("Error creating amqp channel: %s", err)
	}
	defer ch.Close()
	log.Info("Connected to RabbitMQ")

	// Create Session
	s, err := session.New("Bot " + token)
	if err != nil {
		log.Fatal("Failed to create session: %s", err)
	}
	defer s.Close()

	bindEvents(s, ch, &rdb)

	select {}
}

func bindEvents(sess *session.Session, ch *_amqp.Channel, redis *redis.Redis) {
	_handler := &events.Handler{
		EventHandler: handler.EventHandler{
			HandlerR: handler.HandlerR {
				Discord: sess,
				Channel: ch,
				Redis:   redis,
			},
		},
	}
	_handler.Create()

	// Add Handlers
	_handler.AddHandler(_handler.MessageUpdate)
	_handler.AddHandler(_handler.MessageDelete)

	q, err := ch.QueueDeclare(
		"logging_events",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Failed to Declare a Queue: %s", err)
	}

	bindQueues(ch, q.Name)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Failed to Consume a Queue: %s", err)
	}

	go func() {
		for d := range msgs {
			var event handler.Event

			err := json.Unmarshal([]byte(d.Body), &event)

			if err != nil {
				log.Error("Failed to Parse a Message: %s", err)
				d.Nack(false, true)
				continue
			}

			err = _handler.Handle(event)

			if err != nil {
				log.Error("Failed to Handle a Message: %s", err)
				// Disabled because we don't want to constantly loop over the same event
				// d.Nack(false, true)
				continue
			}

			// Ack at the end to signify that we processed this event
			d.Ack(false)
		}
	}()
}

func bindQueues(ch *_amqp.Channel, name string) {
	err := ch.QueueBind(
		name,
		"message.update",
		"events_topic",
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to Bind a Queue: %s", err)
	}
	err = ch.QueueBind(
		name,
		"message.delete",
		"events_topic",
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to Bind a Queue: %s", err)
	}
}
