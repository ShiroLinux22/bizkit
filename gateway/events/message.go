package events

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/streadway/amqp"
)

func (h *EventHandler) MessageCreate(e *gateway.MessageCreateEvent) {
	log.Info("Message: %s", e.Content)
	text := toJson(e)
	log.Info(text)

	err := h.Channel.Publish(
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
