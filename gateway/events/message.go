package events

import (
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/streadway/amqp"
)

func (h *EventHandler) MessageCreate(e *gateway.MessageCreateEvent) {
	log.Info("Message: %s", e.Content)
	text := toJson(e)

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
		log.WarnOnError(err, "Message Not Found")
	}

	log.Info("Deleted: %s", m.Content)
	text := toJson(m)

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