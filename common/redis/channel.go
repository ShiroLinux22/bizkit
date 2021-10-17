package redis

import (
	"encoding/json"
	"fmt"

	"github.com/chakernet/ryuko/common/util"
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
		return cmd.Err();
	}

	return nil
}

func (r *Redis) GetChannel(id discord.ChannelID) (*discord.Channel, error)  {
	key := fmt.Sprintf("channels:%s", id.String())
	cmd := r.Client.Do(ctx, "JSON.GET", key, ".")

	if cmd.Err() != nil {
		if cmd.Err().Error() == "redis: nil" {
			return nil, nil
		}
		return nil, cmd.Err();
	}

	text, err := cmd.Text()
	if err != nil {
		return nil, err;
	}
	bytes := []byte(text)

	var channel discord.Channel

	json.Unmarshal(bytes, &channel)

	return nil, nil
}

func (r *Redis) DeleteChannel(id discord.ChannelID) error  {
	key := fmt.Sprintf("channels:%s", id.String())
	cmd := r.Client.Do(ctx, "JSON.DEL", key, ".")

	if cmd.Err() != nil {
		return cmd.Err();
	}

	return nil
}