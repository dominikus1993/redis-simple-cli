package redisdb

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(ctx context.Context, addr string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "", // no password set
		DB:           0,  // use default DB
		DialTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
		ReadTimeout:  20 * time.Minute,
		PoolTimeout:  20 * time.Minute,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("%w, cannot connect to redis", err)
	}
	return &RedisClient{Client: client}, nil
}

func (client *RedisClient) Close(ctx context.Context) {
	err := client.Client.Close()
	if err != nil {
		panic(err)
	}
}
