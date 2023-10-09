package redisfoo

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func connect(url, password string, DB int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       DB,
	})
	if rdb == nil {
		panic(rdb)
	}
	return rdb
}

func connectTest(url, password string, DB int) string {
	rdb := connect(url, password, DB)
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

func zaddFoo(url, password string, DB int) {
	rdb := connect(url, password, DB)
	r, e := rdb.ZAdd(context.Background(), "LGR_S1_TEST", &redis.Z{
		Score: 124, Member: 123,
	}).Result()
	fmt.Printf("r = %v, e = %v\n", r, e)
}

func getFoo(url, password string, DB int) {
	rdb := connect(url, password, DB)
	r, e := rdb.Get(context.Background(), "NON_EXISTS_KEY").Result()
	fmt.Printf("r = %v, e = %v\n", r, e)
}

func hsetFoo(url, password string, DB int) {
	rdb := connect(url, password, DB)
	m := make(map[string]interface{})
	m["field"] = "value"
	r, e := rdb.HSet(context.Background(), "HSET_TEST", m).Result()
	fmt.Printf("r = %v, e = %v\n", r, e)
}
