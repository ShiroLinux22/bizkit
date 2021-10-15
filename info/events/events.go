package events

import (
	"encoding/json"

	"github.com/chakernet/ryuko/info/util"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

var (
	log = util.Logger {
		Name: "EventHandler",
	}
)

type EventHandler struct {
	Discord *session.Session
}

type Event struct {
	Type string `json:"type"`
	Shard int `json:"shard,omitempty"`
	Data string `json:"data,omitempty"`
}

func (h *EventHandler) Handle(e Event) {
	raw := []byte(e.Data)

	switch e.Type {
	case "MESSAGE_CREATE":
		var data discord.Message

		err := json.Unmarshal(raw, &data)

		if err != nil {
			log.Error("Error trying to parse Event: %s", err)
			return
		}

		mem, err := h.Discord.Member(data.GuildID, data.Author.ID)

		if err != nil {
			log.Error("Error trying to fetch member: %s", err)
			return
		}

		h.MessageCreate(&gateway.MessageCreateEvent{
			Message: data,
			Member: mem,
		})
		break;

	default:
		log.Error("Unknown Event: %s", e.Type)
	}
}