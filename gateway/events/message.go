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

	err = h.Channel.Publish(
		"events_topic",
		"message.create",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(text),
		},
	)
	
	if err != nil {
		log.Error("Failed to send to RabbitMQ", err)
	}
}
