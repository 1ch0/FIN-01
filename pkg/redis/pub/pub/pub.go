package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func pub() {
	ctx := context.Background()
	rdbPub := redis.NewClient(&redis.Options{
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		// default: localhost:6379
		Addr: "139.198.166.89:56379",
	})
	intCmd := rdbPub.Publish(ctx, "example:channel", "明天放假")
	fmt.Println(intCmd.Result())
}

func main() {
	pub()
}
