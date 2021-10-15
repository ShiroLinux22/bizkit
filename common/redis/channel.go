package redis

import (
	"encoding/json"
	"fmt"

	"github.com/chakernet/ryuko/common/util"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/go-redis/redis/v8"
)

func SetChannel(conn *redis.Client, c *discord.Channel) error {
	key := fmt.Sprintf("channels:%s", c.ID.String())

	data, err := util.ToJson(c)
	if err != nil {
		return err
	}

	cmd := conn.Do(ctx, "JSON.SET", key, ".", data)
	if cmd.Err() != nil {
		return cmd.Err();
	}

	return nil
}

func GetChannel(conn *redis.Client, id discord.ChannelID) (discord.Channel, error)  {
	key := fmt.Sprintf("channels:%s", id.String())
	cmd := conn.Do(ctx, "JSON.GET", key, ".")

	if cmd.Err() != nil {
		return discord.Channel{}, cmd.Err();
	}

	text, err := cmd.Text()
	if err != nil {
		return discord.Channel{}, err;
	}
	bytes := []byte(text)

	var channel discord.Channel

	json.Unmarshal(bytes, &channel)

	return channel, nil
}