package redis

import (
	"github.com/atompi/metrics-post-station/pkg/options"
	"github.com/redis/go-redis/v9"
)

func New(opts options.RedisOptions) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     opts.Addr,
		Password: opts.Password,
		DB:       opts.DB,
	})
}
