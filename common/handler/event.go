package handler

import (
	"encoding/json"
	"errors"

	"github.com/diamondburned/arikawa/v3/session"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
)

type IEventHandler interface {
	MessageCreate(*gateway.MessageCreateEvent) error;
}

type EventHandler struct {
	IEventHandler

	Channel *amqp.Channel
	Discord *session.Session
	Redis *redis.Client
}

// R means 'reduced'
type EventHandlerR struct {
	Channel *amqp.Channel
	Discord *session.Session
	Redis *redis.Client
}

type Event struct {
	Type string `json:"type"`
	Shard int `json:"shard,omitempty"`
	Data string `json:"data,omitempty"`
}

func (h *EventHandler) Handle(e Event) error {
	raw := []byte(e.Data)

	switch e.Type {
	case "MESSAGE_CREATE":
		var data discord.Message

		err := json.Unmarshal(raw, &data)

		if err != nil {
			return err
		}

		mem, err := h.Discord.Member(data.GuildID, data.Author.ID)

		if err != nil {
			return err
		}

		err = h.MessageCreate(&gateway.MessageCreateEvent{
			Message: data,
			Member: mem,
		})

		if err != nil {
			return err
		}
		break;

	default:
		return errors.New("invalid event")
	}

	return nil
}