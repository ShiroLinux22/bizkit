/*
	Handler(s) for message-related events
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
	"fmt"
	"os"
	"strings"

	"github.com/diamondburned/arikawa/v3/gateway"
)

func (h *Handler) MessageCreate(m *gateway.MessageCreateEvent) error {
	if m.Author.Bot {
		return nil
	}

	if !strings.HasPrefix(m.Content, os.Getenv("BOT_PREFIX")) {
		return nil
	}

	h.Discord.SendTextReply(m.ChannelID, "LMAO", m.ID)

	return nil
}

func (h *Handler) MessageUpdate(m *gateway.MessageUpdateEvent) error {
	fmt.Printf("content: %s\n", m.Content)

	return nil
}

func (h *Handler) MessageDelete(m *gateway.MessageDeleteEvent) error {
	fmt.Printf("content: %s\n", )

	return nil
}