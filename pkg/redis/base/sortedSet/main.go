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
		DB:   2,
	})

	// 向有序集合添加元素, 并返回成功添加的个数
	intCmd := rdb.ZAdd(ctx, "sorted:set1", &redis.Z{Score: 10, Member: "s1"}, &redis.Z{Score: 20, Member: "s2"})
	intCmd = rdb.ZAdd(ctx, "sorted:set2", &redis.Z{Score: 10, Member: "e1"}, &redis.Z{Score: 20, Member: "e2"})
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 获取有序集合的元素个数
	intCmd = rdb.ZCard(ctx, "sorted:set1")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 获取成员分数值
	floatCmd := rdb.ZScore(ctx, "sorted:set1", "s2")
	fmt.Println(floatCmd.Result())
	// 20 <nil>

	// 获取有序集合中指定分数段中元素的个数
	intCmd = rdb.ZCount(ctx, "sorted:set1", "15", "30")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 获取指定字典区间内成员数量, 包含s1及之后的所有成员
	intCmd = rdb.ZLexCount(ctx, "sorted:set1", "[s1", "-")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 获取指定索引区间内的成员
	strSliceCmd := rdb.ZRange(ctx, "sorted:set1", 0, -1)
	fmt.Println(strSliceCmd.Result())
	// [s1 s2] <nil>

	// 获取指定索引区间内带有排序分数的成员
	zSliceCmd := rdb.ZRangeWithScores(ctx, "sorted:set1", 0, -1)
	fmt.Println(zSliceCmd.Result())
	// [{10 s1} {20 s2}] <nil>

	// 获取指定分数区间内的成员
	strSliceCmd = rdb.ZRangeByScore(ctx, "sorted:set1", &redis.ZRangeBy{Max: "10"})
	fmt.Println(strSliceCmd.Result())
	// [s1] <nil>

	// 通过字典区间返回有序集合的成员。直接对成员值作区间操作, (->不包含，[->包含
	strSliceCmd = rdb.ZRangeByLex(ctx, "sorted:set1", &redis.ZRangeBy{Min: "-", Max: "(s2"})
	fmt.Println(strSliceCmd.Result())
	// [s1] <nil>
	strSliceCmd = rdb.ZRangeByLex(ctx, "sorted:set1", &redis.ZRangeBy{Min: "-", Max: "[s2"})
	fmt.Println(strSliceCmd.Result())
	// [s1 s2] <nil>
	strSliceCmd = rdb.ZRangeByLex(ctx, "sorted:set1", &redis.ZRangeBy{Min: "(s1", Max: "[s2"})
	fmt.Println(strSliceCmd.Result())
	// [s2] <nil>
	strSliceCmd = rdb.ZRangeByLex(ctx, "sorted:set1", &redis.ZRangeBy{Min: "[s1", Max: "[s2"})
	fmt.Println(strSliceCmd.Result())
	// [s1 s2] <nil>

	// 获取指定成员的索引
	intCmd = rdb.ZRank(ctx, "sorted:set1", "s2")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 增加指定成员的分数
	floatCmd = rdb.ZIncr(ctx, "sorted:set1", &redis.Z{Score: 40, Member: "s2"})
	fmt.Println(floatCmd.Result())
	// 60 <nil>

	// 增加指定成员的分数
	floatCmd = rdb.ZIncrBy(ctx, "sorted:set1", 0.5, "s1")
	fmt.Println(floatCmd.Result())
	// 10.5 <nil>

	// 将交集存储在新的有序集合中
	intCmd = rdb.ZInterStore(ctx, "zStoreInter", &redis.ZStore{Keys: []string{"sorted:set1", "sorted:set2"}})
	fmt.Println(intCmd.Result())
	// 0 <nil>
	fmt.Println(rdb.ZRange(ctx, "zStoreInter", 0, -1).Result())
	// [] <nil>

	// 将并集存储在新的有序集中
	intCmd = rdb.ZUnionStore(ctx, "zStoreUnion", &redis.ZStore{Keys: []string{"sorted:set1", "sorted:set2"}})
	fmt.Println(intCmd.Result())
	// 4 <nil>
	fmt.Println(rdb.ZRange(ctx, "zStoreUnion", 0, -1).Result())
	// [e1 s1 e2 s2] <nil>

	// 返回索引区间内的成员，通过分数从高到低排列
	strSliceCmd = rdb.ZRevRange(ctx, "zStoreUnion", 0, -1)
	fmt.Println(strSliceCmd.Result())
	// [s2 e2 s1 e1] <nil>

	// 返回指定分数区间内的成员，通过分数从高到低排列
	strSliceCmd = rdb.ZRevRangeByScore(ctx, "zStoreUnion", &redis.ZRangeBy{Max: "20"})
	fmt.Println(strSliceCmd.Result())
	// [e2 s1 e1] <nil>

	// 按分数从高到低排列，返回指定成员排名
	intCmd = rdb.ZRevRank(ctx, "zStoreUnion", "s1")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 移除成员
	intCmd = rdb.ZRem(ctx, "zStoreUnion", "s1")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 移除指定分数段内的成员
	intCmd = rdb.ZRemRangeByScore(ctx, "zStoreUnion", "0", "25")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 移除指定排名区间内的成员
	intCmd = rdb.ZRemRangeByRank(ctx, "zStoreUnion", 1, 1)
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 移除指定字典区间内的成员
	intCmd = rdb.ZRemRangeByLex(ctx, "zStoreUnion", "[e1", "(e2")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 扫描返回值匹配e*的1个元素
	scanCmd := rdb.ZScan(ctx, "sorted:set2", 0, "e*", 1)
	var res []string
	res, cursor, err := scanCmd.Result()
	fmt.Println(res, cursor, err)
	// [e1 10 e2 20] 0 <nil>

	// 继续从上次游标处开始扫描
	scanCmd = rdb.ZScan(ctx, "sorted:set2", cursor, "e*", 1)
	fmt.Println(scanCmd.Result())
	// [e1 10 e2 20] 0 <nil>
}
