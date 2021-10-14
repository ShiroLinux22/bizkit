package events

import (
	"encoding/json"

	"github.com/chakernet/ryuko/gateway/util"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/streadway/amqp"
)

var (
	log = util.Logger {
		Name: "EventHandler",
	}
)

type EventHandler struct {
	Channel *amqp.Channel
	Discord *session.Session
}

func toJson(i ...interface {}) (string) {
	formatted, _ := json.Marshal(i)
	text := string(formatted)

	return text
}