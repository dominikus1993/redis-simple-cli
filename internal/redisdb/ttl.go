package redisdb

import "context"

func (redis *RedisClient) removeKeysWithoutTTL(ctx context.Context, keys []string) error {
	for _, key := range keys {
		_, err := redis.client.TTL(ctx, key).Result()
		if err != nil {
			return err
		}

	}
	return nil
}
