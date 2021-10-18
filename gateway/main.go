/*
	Main file for the gateway module
    Copyright (C) 2021 jacany <jack@chaker.net>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"context"
	"os"

	"github.com/chakernet/ryuko/common/amqp"
	"github.com/chakernet/ryuko/common/handler"
	"github.com/chakernet/ryuko/common/redis"
	"github.com/chakernet/ryuko/common/util"
	"github.com/chakernet/ryuko/gateway/events"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/joho/godotenv"
	_amqp "github.com/streadway/amqp"
)

func main() {
	log := util.Logger {
		Name: "main",
	}

	log.Info("Initializing...")
	// Load Env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Failed to load env: ", err)
	}
	token := os.Getenv("BOT_TOKEN")
	
	// Connect to Redis
	rdb := redis.Redis {}
	client := rdb.Connect()
	defer client.Close()
	log.Info("Connected to Redis")

	// Connect to rabbitmq
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
		log.Fatal("Failed to create session: ", err)
	}

	bindEvents(s, ch, &log, &rdb)

	// Add the needed Gateway intents.
	s.AddIntents(gateway.IntentGuilds)
	s.AddIntents(gateway.IntentGuildMessages)
	s.AddIntents(gateway.IntentGuildMembers)
	s.AddIntents(gateway.IntentGuildBans)
	s.AddIntents(gateway.IntentGuildInvites)
	s.AddIntents(gateway.IntentGuildVoiceStates)
	s.AddIntents(gateway.IntentDirectMessageReactions)

	// Open a new Session for events
	err = s.Open(context.Background())
	if err != nil {
		log.Fatal("Failed to create connection: ", err)
	}
	defer s.Close()

	// Get bot user
	u, err := s.Me()
	if err != nil {
		log.Fatal("Failed to get user: ", err)
	}
	log.Info("Started as %s", u.Username)

	// Block forever.
	select {}
}

func bindEvents(s *session.Session, ch *_amqp.Channel, log *util.Logger, rdb *redis.Redis) {
	err := ch.ExchangeDeclare(
		"events_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare an exchange: ", err)
	}

	handler := events.Handler {
		EventHandlerR: handler.EventHandlerR{
			Discord: s,
			Channel: ch,
			Redis: rdb,
		},
	}

	// Message Events (GUILD_MESSAGES & DIRECT_MESSAGES)
	s.AddHandler(handler.MessageCreate)
	s.AddHandler(handler.MessageUpdate)
	s.AddHandler(handler.MessageDelete)

	// Channel Events (GUILDS)
	s.AddHandler(handler.ChannelCreate)
	s.AddHandler(handler.ChannelUpdate)
	s.AddHandler(handler.ChannelDelete)
}