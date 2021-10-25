/*
	Interaction Handler Struct(s)
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
	"errors"
	"reflect"
	"sync"

	"github.com/diamondburned/arikawa/v3/discord"
)

type interaction struct {
	Type discord.InteractionType
	fullName reflect.Type

	callback reflect.Value
}

type InteractionHandler struct {
	HandlerR

	mutex sync.RWMutex
	mods map[string]interaction
}

func (h *InteractionHandler) Create() {
	h.mods = make(map[string]interaction)
}

func (h *InteractionHandler) Call(ev interface{}) error {
	evV := reflect.ValueOf(ev)
	evT := evV.Type()

	h.mutex.RLock()
	defer h.mutex.RUnlock()

	for _, value := range h.mods {
		if value.fullName != evT {
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

func (h *InteractionHandler) AddHandler(fn interface{}) (error) {
	r, err := newInteraction(fn)
	if err != nil {
		return err
	}

	h.mutex.Lock()
	h.mods[r.fullName.String()] = r
	h.mutex.Unlock()

	return nil
}

func newInteraction(unknown interface{}) (interaction, error) {
	fnV := reflect.ValueOf(unknown)
	fnT := fnV.Type()

	interaction := interaction {
		callback: fnV,
	}

	if fnT.Kind() != reflect.Func {
		return interaction, errors.New("only functions are accepted")
	}

	if fnT.NumIn() != 1 {
		return interaction, errors.New("function can only accept 1 argument")
	}

	interaction.fullName = fnT.In(0)

	kind := interaction.fullName.Kind()

	if kind != reflect.Ptr && kind != reflect.Interface {
		return interaction, errors.New("function argument is not a pointer")
	}

	return interaction, nil
}