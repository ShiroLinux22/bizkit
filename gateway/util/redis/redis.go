package redis

import (
	"context"
	"os"
	"time"

	"github.com/chakernet/ryuko/gateway/util"
	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
	log = util.Logger {
		Name: "redis",
	}
	dur, err = time.ParseDuration("4h")
)


func Connect() *redis.Client {
	uri := os.Getenv("REDIS_URI")
	auth := os.Getenv("REDIS_AUTH")

	client := redis.NewClient(&redis.Options{
		Addr: uri,
		Password: auth,
		DB: 0,
	})

	log.Info("Connected to Redis")

	return client
}