package main

import (
	"os"

	"github.com/chakernet/ryuko/info/events"
	"github.com/chakernet/ryuko/info/util"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/joho/godotenv"
)

func main() {
	log := util.Logger {
		Name: "main",
	}
	
	// Load Env
	err := godotenv.Load("../.env")
	log.FatalOnError(err, "Failed to load env")
	token := os.Getenv("BOT_TOKEN")

	// Create Session
	s, err := session.New("Bot " + token)
	log.FatalOnError(err, "Failed to create session")

	bindEvents(s)

	select {}
}

func bindEvents(sess *session.Session) {
	handler := events.EventHandler {
		Discord: sess,
	}

	handler.Handle(events.Event{
		Type: "MESSAGE_CREATE",
		Shard: 0,
		Data: `{"id":"898568570668195930","channel_id":"867791409008607276","guild_id":"867791409008607273","type":0,"flags":0,"tts":false,"pinned":false,"mention_everyone":false,"mentions":[],"mention_roles":[],"author":{"id":"739969527957422202","username":"jacany","discriminator":"0001","avatar":"29cc2fc918ff8475eeae071e06ba39fd","public_flags":64},"content":"f","timestamp":"2021-10-15T13:50:41Z","edited_timestamp":null,"attachments":[],"embeds":[],"nonce":"898568569896304640","member":{"user":{"id":null,"username":"","discriminator":"","avatar":""},"roles":["889865510232162354"],"joined_at":"2021-07-22T15:33:14Z","premium_since":null,"deaf":false,"mute":false,"pending":false}}`,
	})
}