package v1

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func Get(c *redis.Client, key string, timeout int) (res string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	res, err = c.Get(ctx, key).Result()
	cancel()
	return
}

func Set(c *redis.Client, key string, value string, timeout int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	err := c.Set(ctx, key, value, 60*time.Second).Err()
	cancel()
	return err
}
