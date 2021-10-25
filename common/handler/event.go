/*
	Event Handler Struct(s)
    Copyright (C) 2021 Jack C <jack@chaker.net>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package handler

import (
	"encoding/json"
	"errors"
	"reflect"
	"sync"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
)

// Event Structs
type MessageCreateEvent struct {
	Member *discord.Member `json:"member"`
	*discord.Message `json:"message"`
}
type MessageUpdateEvent struct {
	Member *discord.Member `json:"member,omitempty"`
	Before *discord.Message `json:"before,omitempty"`
	After *discord.Message `json:"after"`
}

type MessageDeleteEvent struct {
	Member *discord.Member `json:"member"`
	*discord.Message `json:"message,omitempty"`
}


// Other stuff
type event struct {
	event reflect.Type
	callback reflect.Value
}

type EventHandler struct {
	HandlerR

	mutex sync.RWMutex
	mods map[string]event
}

type Event struct {
	Type string `json:"type"`
	Data string `json:"data,omitempty"`
}

func (h *EventHandler) Create() {
	h.mods = make(map[string]event)
}

func (h *EventHandler) Call(ev interface{}) error {
	evV := reflect.ValueOf(ev)
	evT := evV.Type()

	h.mutex.RLock()
	defer h.mutex.RUnlock()

	for _, value := range h.mods {
		if value.event != evT {
			continue
		}
		err := make(chan error, 1)
		go func() {
			fun := value.callback.Call([]reflect.Value{evV})
			
			if reflect.TypeOf(fun[0]).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
				er, _ := fun[0].Interface().(error)
				err <- er
				return
			}

			err <- nil
			return
		}()

		if err != nil {
			return <-err
		}
	}

	return nil
}

func (h *EventHandler) AddHandler(fn interface{}) (error) {
	r, err := newEvent(fn)
	if err != nil {
		return err
	}

	h.mutex.Lock()
	h.mods[r.event.String()] = r
	h.mutex.Unlock()

	return nil
}

func newEvent(unknown interface{}) (event, error) {
	fnV := reflect.ValueOf(unknown)
	fnT := fnV.Type()

	handler := event {
		callback: fnV,
	}

	if fnT.Kind() != reflect.Func {
		return handler, errors.New("only functions are accepted")
	}

	if fnT.NumIn() != 1 {
		return handler, errors.New("function can only accept 1 argument")
	}

	handler.event = fnT.In(0)

	kind := handler.event.Kind()

	if kind != reflect.Ptr && kind != reflect.Interface {
		return handler, errors.New("function argument is not a pointer")
	}

	return handler, nil
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

		err = h.Call(&gateway.MessageCreateEvent{
			Message: data,
			Member:  mem,
		})

		if err != nil {
			return err
		}
		break

	case "MESSAGE_UPDATE":
		var data MessageUpdateEvent

		err := json.Unmarshal(raw, &data)

		if err != nil {
			return err
		}

		err = h.Call(&data)
		
		if err != nil {
			return err
		}
		break

	case "MESSAGE_DELETE":
		var data MessageDeleteEvent

		err := json.Unmarshal(raw, &data)

		if err != nil {
			return err
		}

		err = h.Call(&data)
		
		if err != nil {
			return err
		}
		break

	default:
		return errors.New("invalid event")
	}

	return nil
}
