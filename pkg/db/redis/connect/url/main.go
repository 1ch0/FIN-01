package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	options, err := redis.ParseURL("redis://:eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81@139.198.166.89:56379/1")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(options)
	// ping一下检查是否连通
	pingResult, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	// PONG
	fmt.Println(pingResult)
	defer rdb.Close()

	rdb.Ping(context.Background())
}
