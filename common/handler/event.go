/*
	Event Handler Struct(s)
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

package handler

import (
	"encoding/json"
	"errors"

	"github.com/chakernet/ryuko/common/redis"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/streadway/amqp"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
)

type IEventHandler interface {
	// Message Events (GUILD_MESSAGES & DIRECT_MESSAGES)
	MessageCreate(*gateway.MessageCreateEvent) error
	MessageUpdate(*gateway.MessageUpdateEvent) error
	MessageDelete(*gateway.MessageDeleteEvent) error
	MessageDeleteBulk(*gateway.MessageDeleteBulkEvent) error // guild only

	// Role Events (GUILDS)
	GuildRoleCreate(*gateway.GuildRoleCreateEvent) error
	GuildRoleUpdate(*gateway.GuildRoleUpdateEvent) error
	GuildRoleDelete(*gateway.GuildRoleDeleteEvent) error

	// Guild Events (GUILDS)
	GuildCreate(*gateway.GuildCreateEvent) error
	GuildUpdate(*gateway.GuildUpdateEvent) error
	GuildDelete(*gateway.GuildDeleteEvent) error

	// Channel Events (GUILDS)
	ChannelCreate(*gateway.ChannelCreateEvent) error
	ChannelUpdate(*gateway.ChannelUpdateEvent) error
	ChannelDelete(*gateway.ChannelDeleteEvent) error

	// Member Events (GUILD_MEMBERS)
	GuildMemberAdd(*gateway.GuildMemberAddEvent) error
	GuildMemberUpdate(*gateway.GuildMemberUpdateEvent) error
	GuildMemberRemove(*gateway.GuildMemberRemoveEvent) error

	// Ban Events (GUILD_BANS)
	GuildBanAdd(*gateway.GuildBanAddEvent) error
	GuildBanRemove(*gateway.GuildBanRemoveEvent) error

	// Invite Events (GUILD_INVITES)
	InviteCreate(*gateway.InviteCreateEvent) error
	InviteDelete(*gateway.InviteDeleteEvent) error

	// Voice State Events (GUILD_VOICE_STATES)
	VoiceStateUpdate(*gateway.VoiceStateUpdateEvent) error

	// Guild Message Reactions Events (GUILD_MESSAGE_REACTIONS & DIRECT_MESSAGE_REACTIONS)
	MessageReactionAdd(*gateway.MessageReactionAddEvent) error
	MessageReactionRemove(*gateway.MessageReactionRemoveEvent) error
	MessageReactionRemoveAll(*gateway.MessageReactionRemoveAllEvent) error
	MessageReactionRemoveEmoji(*gateway.MessageReactionRemoveEmojiEvent) error
}

type EventHandler struct {
	IEventHandler

	Channel *amqp.Channel
	Discord *session.Session
	Redis   *redis.Redis
}

// R means 'reduced'
type EventHandlerR struct {
	Channel *amqp.Channel
	Discord *session.Session
	Redis   *redis.Redis
}

type Event struct {
	Type string `json:"type"`
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

		mem, err := h.Redis.GetMember(data.GuildID, data.Author.ID)
		if mem == nil && err == nil {
			mem, err = h.Discord.Member(data.GuildID, data.Author.ID)

			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		err = h.MessageCreate(&gateway.MessageCreateEvent{
			Message: data,
			Member:  mem,
		})

		if err != nil {
			return err
		}
		break

	case "MESSAGE_UPDATE":
		var data discord.Message

		err := json.Unmarshal(raw, &data)

		if err != nil {
			return err
		}

		mem, err := h.Redis.GetMember(data.GuildID, data.Author.ID)
		if mem == nil && err == nil {
			mem, err = h.Discord.Member(data.GuildID, data.Author.ID)

			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		err = h.MessageUpdate(&gateway.MessageUpdateEvent{
			Message: data,
			Member:  mem,
		})

		if err != nil {
			return err
		}
		break

	case "MESSAGE_DELETE":
		break

	default:
		return errors.New("invalid event")
	}

	return nil
}
