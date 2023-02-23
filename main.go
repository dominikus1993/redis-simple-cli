package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/dominikus1993/redis-simple-cli/internal/redisdb"
	"github.com/urfave/cli/v2"
)

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

			n, err := redis.ScanAndRemoveKeysWithoutTTL(ctx.Context)
			if err != nil {
				cli.Exit(err, 1)
			}
			return cli.Exit(fmt.Sprintf("processed keys %d", n), 0)
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
