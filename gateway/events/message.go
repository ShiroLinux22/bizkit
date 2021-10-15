package events

import (
	"github.com/chakernet/ryuko/gateway/util"
	"github.com/chakernet/ryuko/gateway/util/redis"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/streadway/amqp"
)

func (h *EventHandler) MessageCreate(e *gateway.MessageCreateEvent) {
	log.Info("Message: %s", e.Content)
	cha, err := h.Discord.Channel(e.ChannelID)
	if err != nil {
		log.Error("Failed to get channel:", err)
		return
	}
	redis.SetChannel(h.Redis, cha)
	text := util.ToJson(e)

	h.Channel.Publish(
		"events_topic",
		"message.create",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(text),
		},
	)
}

func (h *EventHandler) MessageDelete(e *gateway.MessageDeleteEvent) {
	// Get Message
	m, err := h.Discord.Message(e.ChannelID, e.ID)
	if err != nil {
		log.Error("Failed to get message:", err)

		return
	}

	log.Info("Deleted: %s", m.Content)
	text := util.ToJson(m)

	h.Channel.Publish(
		"events_topic",
		"message.delete",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(text),
		},
	)
}