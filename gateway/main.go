package main

import (
	"context"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/joho/godotenv"
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

	// Create Session
	s, err := session.New("Bot " + token)
	failOnError(err, "Failed to create session")

	s.AddHandler(func(c *gateway.MessageCreateEvent) {
		log.Println(c.Author.Username, "sent", c.Content)
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