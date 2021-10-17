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