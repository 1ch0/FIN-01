package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		Addr:     "139.198.166.89:56379",
	})

	// 为已存在的列表添加值， 不存在时返回length为0， 且push不会成功,也不会打印错误
	intCmd := rdb.RPushX(ctx, "not:exists:list", "a", "b", "c")
	length, err := intCmd.Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(length, err)
	// 0 <nil>

	// 向列表添加元素，无论列表是否存在
	intCmd = rdb.RPush(ctx, "example:list", "a", "b", "c")
	length, err = intCmd.Result()
	fmt.Println(length, err)
	// 3 <nil>

	// 移除列表的最后一个元素
	strCmd := rdb.RPop(ctx, "example:list")
	result, err := strCmd.Result()
	fmt.Println(result, err)
	// c <nil>

	// 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	strCmd = rdb.RPopLPush(ctx, "example:list", "example:list2")
	result, err = strCmd.Result()
	fmt.Println(result, err)
	// b <nil>

	// 列出长度
	intCmd = rdb.LLen(ctx, "example:list2")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 列出某个索引位置的值
	strCmd = rdb.LIndex(ctx, "example:list2", 0)
	fmt.Println(strCmd.Result())
	// b <nil>

	// 列出某个范围索引的值
	strSliceCmd := rdb.LRange(ctx, "example:list2", 0, -1)
	fmt.Println(strSliceCmd.Result())
	// [b] <nil>

	// 弹出最后一个元素，如果没有则会阻塞列表直到等待超时或发现可弹出元素为止。
	// 可以有多个key，当第一个key中没有值时，才会从第二个key中pop
	strSliceCmd = rdb.BRPop(ctx, 5*time.Second, "example:list", "example:list2")
	var results []string
	results, err = strSliceCmd.Result()
	fmt.Println(results, err)
	// [example:list a] <nil>

	// 将a插入到列表中的b之前, 然后返回插入后列表元素
	intCmd = rdb.LInsert(ctx, "example:list2", "before", "b", "a")
	// 等价于
	// intCmd = rdb.LInsertBefore(ctx, "example:list2", "b", "a")
	fmt.Println(intCmd.Result())
	// 2 <nil>
	fmt.Println(rdb.LRange(ctx, "example:list2", 0, -1).Result())
	// [a b] <nil>

	// 通过索引设置列表元素的值
	statusCmd := rdb.LSet(ctx, "example:list2", 0, "c")
	fmt.Println(statusCmd.Result())
	// OK <nil>

	// 移除元素多少次，0代表全部移除, 返回移除元素的个数
	intCmd = rdb.LRem(ctx, "example:list2", 0, "c")
	fmt.Println(intCmd.Result())
	// 1 <nil>
}
