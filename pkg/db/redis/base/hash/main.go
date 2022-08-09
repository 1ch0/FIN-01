package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		// default: localhost:6379
		Addr: "139.198.166.89:56379",
	})

	// 设置哈希结构的field:value, 三种方式等价
	// HMSet()、HMGet()在redis3已废弃
	intCmd := rdb.HSet(ctx, "example:hash1", "name", "张三", "role", "法外狂徒")
	intCmd = rdb.HSet(ctx, "example:hash2", map[string]interface{}{"name": "李四", "role": "???"})
	intCmd = rdb.HSet(ctx, "example:hash3", []string{"name", "王五", "role", "!!!"})
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 在字段不存在时才设置field:value
	boolCmd := rdb.HSetNX(ctx, "example:hash1", "age", 16)
	fmt.Println(boolCmd.Result())
	// true <nil>

	// 获取哈希字段数量
	intCmd = rdb.HLen(ctx, "example:hash1")
	fmt.Println(intCmd.Result())
	// 3 <nil>

	// 获取指定字段的值
	strCmd := rdb.HGet(ctx, "example:hash1", "role")
	fmt.Println(strCmd.Result())
	// 法外狂徒 <nil>

	// 获取哈希字段
	strSliceCmd := rdb.HKeys(ctx, "example:hash1")
	fmt.Println(strSliceCmd.Result())
	// [name role age] <nil>

	// 获取哈希字段值
	strSliceCmd = rdb.HVals(ctx, "example:hash1")
	fmt.Println(strSliceCmd.Result())
	// [张三 法外狂徒 16] <nil>

	// 获取所有字段和值
	strStrMapCmd := rdb.HGetAll(ctx, "example:hash1")
	fmt.Println(strStrMapCmd.Result())
	// map[age:16 name:张三 role:法外狂徒] <nil>

	// 判断字段是否存在
	boolCmd = rdb.HExists(ctx, "example:hash1", "name")
	fmt.Println(boolCmd.Result())
	// true <nil>

	// 为字段增加值
	intCmd = rdb.HIncrBy(ctx, "example:hash1", "age", 2)
	fmt.Println(intCmd.Result())
	// 18 <nil>

	// 为字段增加浮点数值
	floatCmd := rdb.HIncrByFloat(ctx, "example:hash1", "age", 0.5)
	fmt.Println(floatCmd.Result())
	// 18.5 <nil>

	// 删除字段
	intCmd = rdb.HDel(ctx, "example:hash1", "name", "role", "age")
	fmt.Println(intCmd.Result())
	// 3 <nil>

	// 迭代扫描哈希键值对
	scanCmd := rdb.HScan(ctx, "example:hash2", 0, "name", 1)
	fmt.Println(scanCmd.Result())
	// [name 李四] 0 <nil>
}
