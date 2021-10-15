package redis

import (
	"encoding/json"
	"fmt"

	"github.com/chakernet/ryuko/gateway/util"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/go-redis/redis/v8"
)

func SetChannel(conn *redis.Client, c *discord.Channel) {
	key := fmt.Sprintf("channels:%s", c.ID.String())

	cmd := conn.Do(ctx, "JSON.SET", key, ".", util.ToJson(c))
	if cmd.Err() != nil {
		log.Error("Failed to set channel: ", cmd.Err)
		return;
	}
}

func GetChannel(conn *redis.Client, id discord.ChannelID) discord.Channel  {
	key := fmt.Sprintf("channels:%s", id.String())
	cmd := conn.Do(ctx, "JSON.GET", key, ".")

	if cmd.Err() != nil {
		log.Error("Failed to get channel: ", cmd.Err)
		return discord.Channel{};
	}

	text, err := cmd.Text()
	if err != nil {
		log.Error("Failed to convert to text: ", err)
		return discord.Channel{};
	}
	bytes := []byte(text)

	var channel discord.Channel

	json.Unmarshal(bytes, &channel)

	return channel
}