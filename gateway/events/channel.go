/*
	Handler(s) for channel related events
    Copyright (C) 2021 jacany <jack@chaker.net>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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