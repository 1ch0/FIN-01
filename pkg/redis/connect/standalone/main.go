package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Username: "",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		// default: localhost:6379
		Addr:     "139.198.166.89:56379",
		DB:       1,
		PoolSize: 5,
	})
	// ping一下检查是否连通
	pingResult, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	// PONG
	fmt.Println(pingResult)
}
