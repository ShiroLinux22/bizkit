/*
	Handler(s) for message related events
    Copyright (C) 2021 Jack C <jack@chaker.net>

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
	"github.com/chakernet/bizkit/common/handler"
	"github.com/chakernet/bizkit/common/util"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/streadway/amqp"
)

func (h *Handler) MessageCreate(m *gateway.MessageCreateEvent) {
	// Add Message into Redis
	err := h.Redis.SetMessage(&m.Message)
	if err != nil {
		log.Error("Failed to cache message: %s", err)
		return
	}

	data, err := util.ToJson(m)
	if err != nil {
		log.Error("Failed to parse data to JSON: %s", err)
		return
	}
	event := handler.Event{
		Type: "MESSAGE_CREATE",
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
			Body:        []byte(payload),
		},
	)

	if err != nil {
		log.Error("Failed to send to RabbitMQ: %s", err)
		return
	}
}

func (h *Handler) MessageUpdate(m *gateway.MessageUpdateEvent) {
	message, err := h.Redis.GetMessage(m.ID)
	if err != nil {
		log.Error("Failed to get cached message: %s", err)
	}

	if message == nil {
		log.Info("FUCK")
	}

	// Add Message into Redis
	err = h.Redis.SetMessage(&m.Message)
	if err != nil {
		log.Error("Failed to cache message: %s", err)
		return
	}

	data, err := util.ToJson(handler.MessageUpdateEvent {
		Before: message,
		After: &m.Message,
		Member: m.Member,
	})
	if err != nil {
		log.Error("Failed to parse data to JSON: %s", err)
		return
	}
	event := handler.Event{
		Type: "MESSAGE_UPDATE",
		Data: data,
	}
	payload, err := util.ToJson(event)
	if err != nil {
		log.Error("Failed to parse event to JSON: %s", err)
		return
	}

	err = h.Channel.Publish(
		"events_topic",
		"message.update",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		},
	)

	if err != nil {
		log.Error("Failed to send to RabbitMQ: %s", err)
		return
	}
}

func (h *Handler) MessageDelete(m *gateway.MessageDeleteEvent) {
	cached, err := h.Redis.GetMessage(m.ID)
	if err != nil {
		log.Error("Failed to get cached message: %s", err)
	}

	// Add Message into Redis
	err = h.Redis.DeleteMessage(m.ID)
	if err != nil {
		log.Error("Failed to delete message: %s", err)
		return
	}

	data, err := util.ToJson(cached)
	if err != nil {
		log.Error("Failed to parse data to JSON: %s", err)
		return
	}
	event := handler.Event{
		Type: "MESSAGE_DELETE",
		Data: data,
	}
	payload, err := util.ToJson(event)
	if err != nil {
		log.Error("Failed to parse event to JSON: %s", err)
		return
	}

	err = h.Channel.Publish(
		"events_topic",
		"message.delete",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		},
	)

	if err != nil {
		log.Error("Failed to send to RabbitMQ: %s", err)
		return
	}
}
