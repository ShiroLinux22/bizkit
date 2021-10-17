package events

import (
	"github.com/chakernet/ryuko/common/handler"
	"github.com/chakernet/ryuko/common/util"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/streadway/amqp"
)

func (h *Handler) ChannelCreate(c *gateway.ChannelCreateEvent) {
	err := h.Redis.SetChannel(&c.Channel)
	if err != nil {
		log.Error("Failed to cache channel: %s", err)
		return
	}

	data, err := util.ToJson(c)
	if err != nil {
		log.Error("Failed to parse data to JSON: %s", err)
		return
	}
	event := handler.Event {
		Type: "CHANNEL_CREATE",
		Data: data,
	}
	payload, err := util.ToJson(event)
	if err != nil {
		log.Error("Failed to parse event to JSON: %s", err)
		return
	}

	err = h.Channel.Publish(
		"events_topic",
		"channel.create",
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

func (h *Handler) ChannelUpdate(c *gateway.ChannelUpdateEvent) {
	err := h.Redis.SetChannel(&c.Channel)
	if err != nil {
		log.Error("Failed to cache channel: %s", err)
		return
	}

	data, err := util.ToJson(c)
	if err != nil {
		log.Error("Failed to parse data to JSON: %s", err)
		return
	}
	event := handler.Event {
		Type: "CHANNEL_UPDATE",
		Data: data,
	}
	payload, err := util.ToJson(event)
	if err != nil {
		log.Error("Failed to parse event to JSON: %s", err)
		return
	}

	err = h.Channel.Publish(
		"events_topic",
		"channel.update",
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

func (h *Handler) ChannelDelete(c *gateway.ChannelDeleteEvent) {
	chann, err := h.Redis.GetChannel(c.ID)
	if chann == nil && err == nil {
		chann, err = h.Discord.Channel(c.ID)

		if err != nil {
			return 
		}
	} else if err != nil {
		return 
	}

	err = h.Redis.DeleteChannel(c.ID)
	if err != nil {
		log.Error("Failed to delete cached channel: %s", err)
		return
	}

	data, err := util.ToJson(chann)
	if err != nil {
		log.Error("Failed to parse data to JSON: %s", err)
		return
	}
	event := handler.Event {
		Type: "CHANNEL_DELETE",
		Data: data,
	}
	payload, err := util.ToJson(event)
	if err != nil {
		log.Error("Failed to parse event to JSON: %s", err)
		return
	}

	err = h.Channel.Publish(
		"events_topic",
		"channel.delete",
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