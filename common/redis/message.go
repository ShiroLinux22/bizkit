package redis

import (
	"encoding/json"
	"fmt"

	"github.com/chakernet/ryuko/common/util"
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
		return cmd.Err();
	}

	return nil
}

func (r *Redis) GetMessage(id discord.MessageID) (*discord.Message, error)  {
	key := fmt.Sprintf("messages:%s", id.String())
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

	var message discord.Message

	json.Unmarshal(bytes, &message)

	return nil, nil
}

func (r *Redis) DeleteMessage(id discord.MessageID) error  {
	key := fmt.Sprintf("messages:%s", id.String())
	cmd := r.Client.Do(ctx, "JSON.DEL", key, ".")

	if cmd.Err() != nil {
		return cmd.Err();
	}

	return nil
}