package events

import (
	"github.com/chakernet/ryuko/common/handler"
	"github.com/chakernet/ryuko/common/redis"
	"github.com/chakernet/ryuko/common/util"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/streadway/amqp"
)

func (h *Handler) MessageCreate(e *gateway.MessageCreateEvent) {
	if e.Author.Bot {
		return
	}
	log.Info("Message: %s", e.Content)
	cha, err := h.Discord.Channel(e.ChannelID)
	if err != nil {
		log.Error("Failed to get channel: %s", err)
		return
	}
	redis.SetChannel(h.Redis, cha)

	data, err := util.ToJson(e)
	if err != nil {
		log.Error("Failed to parse data to JSON: %s", err)
		return
	}
	event := handler.Event {
		Type: "MESSAGE_CREATE",
		Shard: 0,
		Data: data,
	}
	payload, err := util.ToJson(event)
	if err != nil {
		log.Error("Failed to parse event to JSON: %s", err)
		return
	}

	err = h.Channel.Publish(
		"events_topic",
		"message.create",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: []byte(payload),
		},
	)
	
	if err != nil {
		log.Error("Failed to send to RabbitMQ: %s", err)
		return
	}
}
