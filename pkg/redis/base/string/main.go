package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "139.198.166.89:56379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
	})

	statusCmd := rdb.Set(ctx, "language", "golang", 5*time.Second)
	fmt.Println(statusCmd.String())
	// set language golang ex 5: OK
	fmt.Println(statusCmd.Result())
	// OK <nil>

	stringCmd := rdb.Get(ctx, "language")
	fmt.Println(stringCmd.String())
	// get language: golang
	fmt.Println(stringCmd.Result())
	// golang <nil>

	stringCmd2 := rdb.GetSet(ctx, "language", "php")
	fmt.Println(stringCmd2.String())
	// getset language php: golang
	fmt.Println(stringCmd2.Result())
	// golang <nil>

	boolCmd := rdb.SetNX(ctx, "language", "go", 5*time.Second)
	fmt.Println(boolCmd.Result())
	// OK <nil>

	intCmd := rdb.StrLen(ctx, "language")
	fmt.Println(intCmd.Result())
	// 3 <nil>

	intCmd = rdb.Append(ctx, "language", "is the best")
	fmt.Println(intCmd.Result())
	// 14 <nil>
	str, _ := rdb.Get(ctx, "language").Result()
	fmt.Println(len(str))
	// 14

	// statusCmd2 := rdb.MSet(ctx, []interface{}{"php", "world best", "go", 666})
	// statusCmd2 := rdb.MSet(ctx, map[string]interface{}{"php": "world best", "go": 666})
	statusCmd2 := rdb.MSet(ctx, "php", "world best", "go", 666) // 三种方式
	fmt.Println(statusCmd2.Result())
	// OK <nil>

	sliceCmd := rdb.MGet(ctx, "php", "go")
	fmt.Println(sliceCmd.Result())
	// [world best 666] <nil>

	intCmd2 := rdb.Incr(ctx, "go")
	fmt.Println(intCmd2.Result())
	// 667 <nil>

	intCmd = rdb.Decr(ctx, "go")
	fmt.Println(intCmd.Result())
	// 666 <nil>

	intCmd3 := rdb.IncrBy(ctx, "go", 333)
	fmt.Println(intCmd3.Result())
	// 999 <nil>

	intCmd3 = rdb.DecrBy(ctx, "go", 333)
	fmt.Println(intCmd3.Result())
	// 666 <nil>

	floatCmd := rdb.IncrByFloat(ctx, "go", 0.666)
	fmt.Println(floatCmd.Result())
	// 666.666 <nil>
}
