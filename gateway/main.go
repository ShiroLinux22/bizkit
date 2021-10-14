package main

import (
	"context"
	"os"

	"github.com/chakernet/ryuko/gateway/events"
	"github.com/chakernet/ryuko/gateway/util"
	"github.com/chakernet/ryuko/gateway/util/amqp"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/diamondburned/arikawa/v3/utils/handler"
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
	log.FatalOnError(err, "Failed to load env")
	token := os.Getenv("BOT_TOKEN")

	// Connect to rabbitmq
	conn := amqp.Connect()
	defer conn.Close()
	ch := amqp.Channel(conn)
	defer ch.Close()

	// Create Session
	s, err := session.New("Bot " + token)
	log.FatalOnError(err, "Failed to create session")

	s.Handler = handler.New()
	s.Handler.Synchronous = true
	bindEvents(s, ch, &log)

	// Add the needed Gateway intents.
	s.AddIntents(gateway.IntentGuildMessages)

	// Open a new Session for events
	err = s.Open(context.Background())
	log.FatalOnError(err, "Failed to create session")
	defer s.Close()

	// Get bot user
	u, err := s.Me()
	log.FatalOnError(err, "Failed to get bot user")
	log.Info("Started as %s", u.Username)

	// Block forever.
	select {}
}

func bindEvents(s *session.Session, ch *_amqp.Channel, log *util.Logger) {
	err := ch.ExchangeDeclare(
		"events_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	log.FatalOnError(err, "Failed to declare an exchange")

	handler := events.EventHandler {
		Discord: s,
		Channel: ch,
	}

	s.AddHandler(handler.MessageCreate)
}