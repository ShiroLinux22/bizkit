package main

import (
	"encoding/json"
	"os"

	"github.com/chakernet/ryuko/common/amqp"
	"github.com/chakernet/ryuko/common/util"
	"github.com/chakernet/ryuko/info/events"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/joho/godotenv"
	_amqp "github.com/streadway/amqp"
)

var (
	log = util.Logger {
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

	bindEvents(s, ch)

	select {}
}

func bindEvents(sess *session.Session, ch *_amqp.Channel) {
	handler := events.EventHandler {
		Discord: sess,
	}

	q, err := ch.QueueDeclare(
		"info_events",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Failed to Declare a Queue: %s", err)
	}

	err = ch.QueueBind(
		q.Name,
		"message.create",
		"events_topic",
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Failed to Bind a Queue: %s", err)
	}

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
			var parsed events.Event

			err := json.Unmarshal([]byte(d.Body), &parsed)

			if err != nil {
				log.Error("Failed to Parse a Message: %s", err)
				return
			}

			handler.Handle(parsed)

			// Ack at the end to signify that we processed this event
			d.Ack(false)
		}
	}()
}