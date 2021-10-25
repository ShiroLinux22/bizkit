/*
	Caching function(s) for channels
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

package redis

import (
	"encoding/json"
	"fmt"

	"github.com/chakernet/bizkit/common/util"
	"github.com/diamondburned/arikawa/v3/discord"
)

func (r *Redis) SetChannel(c *discord.Channel) error {
	key := fmt.Sprintf("channels:%s", c.ID.String())

	data, err := util.ToJson(c)
	if err != nil {
		return err
	}

	cmd := r.Client.Do(ctx, "JSON.SET", key, ".", data)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (r *Redis) GetChannel(id discord.ChannelID) (*discord.Channel, error) {
	key := fmt.Sprintf("channels:%s", id.String())
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

	var channel discord.Channel

	json.Unmarshal(bytes, &channel)

	return &channel, nil
}

func (r *Redis) DeleteChannel(id discord.ChannelID) error {
	key := fmt.Sprintf("channels:%s", id.String())
	cmd := r.Client.Do(ctx, "JSON.DEL", key, ".")

	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
