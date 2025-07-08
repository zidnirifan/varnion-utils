package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func Ping(ctx context.Context, db *redis.Client) error {
	ping := db.Ping(ctx)
	if ping.Err() != nil {
		return ping.Err()
	}

	return nil
}
