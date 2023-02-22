package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/urfave/cli/v2"
)

func NewRedisClient(ctx context.Context, addr string) (*redis.Client, error) {
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
	return client, nil
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "redis",
				Usage:    "redis",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			redisConnection := ctx.String("redis")
			client, err := NewRedisClient(ctx.Context, redisConnection)
			if err != nil {
				return cli.Exit(errors.Join(err, errors.New("can't connect to redis")), 1)
			}
			client.Scan(ctx.Context)
			return cli.Exit("da", 32)
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
