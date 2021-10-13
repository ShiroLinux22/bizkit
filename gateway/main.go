package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Load Env
	err := godotenv.Load("../.env")
	failOnError(err, "Failed to load env")
	token := os.Getenv("BOT_TOKEN")
	amqpUri := os.Getenv("AMQP_URI")

	// Connect to rabbitmq
	conn, err := amqp.Dial(amqpUri)
	failOnError(err, "Failed to connect to rabbitmq")
	defer conn.Close()
	
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"events",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	// Create Session
	s, err := session.New("Bot " + token)
	failOnError(err, "Failed to create session")

	s.AddHandler(func(c *gateway.MessageCreateEvent) {
		log.Println("LOL")
		formatted, _ := json.Marshal(c)
		text := string(formatted)

		ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body: []byte(text),
			},
		)
	})

	// Add the needed Gateway intents.
	s.AddIntents(gateway.IntentGuildMessages)

	// Open a new Session for events
	err = s.Open(context.Background())
	failOnError(err, "Failed to create session")
	defer s.Close()

	// Get bot user
	u, err := s.Me()
	failOnError(err, "Failed to get bot user")
	log.Println("Started as", u.Username)

	// Block forever.
	select {}
}