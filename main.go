package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/dominikus1993/redis-simple-cli/internal/redisdb"
	"github.com/redis/go-redis/v9"
	"github.com/urfave/cli/v2"
)

type RedisClient struct {
	client *redis.Client
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
			redis, err := redisdb.NewRedisClient(ctx.Context, redisConnection)
			if err != nil {
				return cli.Exit(errors.Join(err, errors.New("can't connect to redis")), 1)
			}
			defer redis.Close(ctx.Context)

			return cli.Exit(fmt.Sprintf("processed keys %d", 10), 0)
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
