package redisdb

import (
	"context"
	"fmt"
	"testing"

	redisc "github.com/dominikus1993/integrationtestcontainers-go/redis"
)

func TestRedisClientConnection(t *testing.T) {
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

	_, err = rdb.Client.Ping(ctx).Result()
	if err != nil {
		t.Fatal(fmt.Errorf("error when tring ping redis server: %w", err))
	}
}
