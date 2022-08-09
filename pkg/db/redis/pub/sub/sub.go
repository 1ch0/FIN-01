// sub.go
package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func sub() {
	ctx := context.Background()

	rdbSub := redis.NewClient(&redis.Options{
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		// default: localhost:6379
		Addr: "139.198.166.89:56379",
	})

	sub := rdbSub.Subscribe(ctx, "example:channel")
	// 匹配pattern进行订阅
	// sub:= rdbSub.PSubscribe(ctx, "*:channel")

	select {
	case message := <-sub.Channel():
		fmt.Println(message.Channel)
		// example:channel
		fmt.Println(message.Pattern)
		//
		fmt.Println(message.Payload)
		// 明天放假了!!!
		fmt.Println(message.PayloadSlice)
		// []
		fmt.Println(message.String())
		// Message<example:channel: 明天放假了!!!>
	}
	defer sub.Close()

	// 获取匹配pattern的所有channels
	strSliceCmd := rdbSub.PubSubChannels(ctx, "*:channel")
	fmt.Println(strSliceCmd.Result())
	// [example:channel] <nil>

	// 获取pattern的个数
	intCmd := rdbSub.PubSubNumPat(ctx)
	fmt.Println(intCmd.Result())
	// 0 <nil>

	// 获取指定channel上的订阅者个数<代码订阅1个+终端命令行订阅1个>
	strIntMapCmd := rdbSub.PubSubNumSub(ctx, "example:channel")
	fmt.Println(strIntMapCmd.Result())
	// map[example:channel:2] <nil>
}

func main() {
	sub()
}
