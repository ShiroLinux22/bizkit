package redis

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

type Redis struct {
	Client *redis.Client
}

func (r *Redis) Connect() *redis.Client {
	uri := os.Getenv("REDIS_URI")
	auth := os.Getenv("REDIS_AUTH")

	client := redis.NewClient(&redis.Options{
		Addr: uri,
		Password: auth,
		DB: 0,
	})

	r.Client = client
	return client
}