package redisdb

import (
	"context"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (redis *RedisClient) RemoveKeysWithoutTTL(ctx context.Context, keys []string) error {
	for _, key := range keys {
		ttl, err := redis.Client.TTL(ctx, key).Result()
		if err != nil {
			return err
		}
		seconds := ttl.Seconds()
		if seconds == -1 {
			_, err := redis.Client.Del(ctx, key).Result()
			if err != nil {
				return errors.Join(err, fmt.Errorf("can't create redis key: %s", key))
			}
		} else if seconds == -2 {
			log.WithField("key", key).Println("Key aleardy exists")
		}

	}
	return nil
}
