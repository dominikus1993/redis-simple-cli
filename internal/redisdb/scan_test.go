package redisdb

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/dominikus1993/go-toolkit/random"
	redisc "github.com/dominikus1993/integrationtestcontainers-go/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestScanAndRemoveKeysWithoutTTLConnection(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	config := redisc.NewRedisContainerConfigurationBuilder().Build()
	// Arrange
	ctx := context.Background()

	container, err := redisc.StartContainer(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	connectionString, err := container.Url(ctx)
	if err != nil {
		t.Fatal(fmt.Errorf("can't download redis conectionstring, %w", err))
	}
	rdb, err := NewRedisClient(ctx, connectionString)
	if err != nil {
		t.Fatal(fmt.Errorf("can't connect to redis, %w", err))
	}
	t.Cleanup(func() {
		rdb.Close(ctx)
	})
	key1 := random.String(10)
	_, err = rdb.Client.Set(ctx, key1, "test", 50000*time.Second).Result()
	if err != nil {
		t.Fatal(fmt.Errorf("error when trying set key1 in redis server: %w", err))
	}
	key2 := random.String(10)
	_, err = rdb.Client.Set(ctx, key2, "test", -1).Result()
	if err != nil {
		t.Fatal(fmt.Errorf("error when trying set key2 in redis server: %w", err))
	}

	n, err := rdb.ScanAndRemoveKeysWithoutTTL(ctx)
	assert.Equal(t, 2, n)
	assert.NoError(t, err)

	key, err := rdb.Client.Get(ctx, key1).Result()
	assert.NoError(t, err)
	assert.Equal(t, "test", key)
	//assert.False(t, errors.Is(err, redis.Nil))

	// Should be deleted because key2 has no expiration
	_, err = rdb.Client.Get(ctx, key2).Result()
	assert.Error(t, err)
	assert.True(t, errors.Is(err, redis.Nil))
}
