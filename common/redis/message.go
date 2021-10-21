/*
	Caching function(s) for messages
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

package redis

import (
	"encoding/json"
	"fmt"

	"github.com/chakernet/bizkit/common/util"
	"github.com/diamondburned/arikawa/v3/discord"
)

func (r *Redis) SetMessage(m *discord.Message) error {
	key := fmt.Sprintf("messages:%s", m.ID.String())

	data, err := util.ToJson(m)
	if err != nil {
		return err
	}

	cmd := r.Client.Do(ctx, "JSON.SET", key, ".", data)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (r *Redis) GetMessage(id discord.MessageID) (*discord.Message, error) {
	key := fmt.Sprintf("messages:%s", id.String())
	cmd := r.Client.Do(ctx, "JSON.GET", key, ".")

	if cmd.Err() != nil {
		if cmd.Err().Error() == "redis: nil" {
			return nil, nil
		}
		return nil, cmd.Err()
	}

	text, err := cmd.Text()
	if err != nil {
		return nil, err
	}
	bytes := []byte(text)

	var message discord.Message

	json.Unmarshal(bytes, &message)

	return &message, nil
}

func (r *Redis) DeleteMessage(id discord.MessageID) error {
	key := fmt.Sprintf("messages:%s", id.String())
	cmd := r.Client.Do(ctx, "JSON.DEL", key, ".")

	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
