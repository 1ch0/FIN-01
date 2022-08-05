package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewRing(&redis.RingOptions{
		Username: "",
		Password: "",
		DB:       0,
		PoolSize: 10,
		Addrs: map[string]string{
			"shard1": ":7000",
			"shard2": ":7001",
			"shard3": ":7002",
		},
	})
	rdb.Ping(ctx)
}
