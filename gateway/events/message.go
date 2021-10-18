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
	"reflect"

	"github.com/chakernet/ryuko/common/handler"
	"github.com/chakernet/ryuko/common/util"
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
	cached, err := h.Redis.GetMessage(m.ID)
	if err != nil {
		log.Error("Failed to get cached message: %s", err)
	}

	// We would like to assume that redis is the source of truth.
	// This *should* be okay to assume in case of some weird
	// edge case where a channel isn't present in local cache
	if util.IsZero(reflect.ValueOf(cached)) {
		cached = &m.Message
	}

	// Add Message into Redis
	err = h.Redis.SetMessage(&m.Message)
	if err != nil {
		log.Error("Failed to cache message: %s", err)
		return
	}

	data, err := util.ToJson(cached)
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
