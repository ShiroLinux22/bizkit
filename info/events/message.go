package events

import (
	"os"
	"strings"

	"github.com/diamondburned/arikawa/v3/gateway"
)

func (h *Handler) MessageCreate(m *gateway.MessageCreateEvent) error {
	if !strings.HasPrefix(m.Content, os.Getenv("BOT_PREFIX")) {
		return nil
	}

	h.Discord.SendTextReply(m.ChannelID, "LMAO", m.ID)

	return nil
}