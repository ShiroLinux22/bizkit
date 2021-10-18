/*
	Struct(s) for handling Redis connections
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
		Addr:     uri,
		Password: auth,
		DB:       0,
	})

	r.Client = client
	return client
}
