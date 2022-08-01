package redisfoo

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func RedisV8MGet(uri string) {
	rcOpts, err := redis.ParseURL(uri)
	if err != nil {
		panic(err)
	}

	rc := redis.NewClient(rcOpts)
	r, err := rc.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("registry redis do error.%v", err.Error()))
	}
	if r != "PONG" {
		panic(fmt.Sprintf("registry redis ping error.%v", r))
	}

	c := rc.HMSet(context.Background(), "test_key", "field1", "value1", "field2", "value2", "field3", "value3")

	pipe := rc.TxPipeline()
	for _, idxKeyName := range newKeys {
		metaResults = append(metaResults, pipe.HMGet(ctx, idxKeyName, "ver", "meta"))
	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
}
