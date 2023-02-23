package redisdb

import (
	"context"
	"errors"
)

func (redis *RedisClient) ScanAndRemoveKeysWithoutTTL(ctx context.Context) (int, error) {
	var cursor uint64
	var n int
	for {
		var keys []string
		var err error
		keys, cursor, err = redis.Client.Scan(ctx, cursor, "*", 10).Result()
		if err != nil {
			return n, errors.Join(err, errors.New("can't scan redis"))
		}
		err = redis.RemoveKeysWithoutTTL(ctx, keys)
		n += len(keys)
		if cursor == 0 {
			break
		}
		if err != nil {
			return n, errors.Join(err, errors.New("can't remove redis keys"))
		}
	}
	return n, nil
}
