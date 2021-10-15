package events

import (
	"github.com/diamondburned/arikawa/v3/gateway"
)

func (h *EventHandler) MessageCreate(m *gateway.MessageCreateEvent) {
	log.Info("Message: %s, Channel: %s, Member: %s", m.Content, m.ChannelID, m.Member.User.Tag())
	//log.Info("Channel: %s, MID: %s", m.ChannelID, m.Message.ID)
	h.Discord.SendTextReply(m.ChannelID, "Fuck you", m.Message.ID)
}