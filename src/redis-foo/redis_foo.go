package redisfoo

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func connectTest(url, password string, DB int) string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       DB,
	})
	if rdb == nil {
		panic(rdb)
	}
	cmd := rdb.Ping(context.Background())
	if cmd == nil {
		panic(cmd)
	}
	result, err := cmd.Result()
	if err != nil {
		panic(err)
	}
	return result
}
