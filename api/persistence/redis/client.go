package redis

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

func New(dsn string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrapf(err, "failed to ping redis server")
	}
	return client, nil
}
