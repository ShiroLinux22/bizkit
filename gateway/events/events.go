package events

import (
	"github.com/chakernet/ryuko/gateway/util"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/go-redis/redis/v8"
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
	Redis *redis.Client
}
