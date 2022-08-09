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

	// 向集合添加元素, 返回成功添加的元素个数
	intCmd := rdb.SAdd(ctx, "example:set1", "s1", "s2", "s3")
	intCmd = rdb.SAdd(ctx, "example:set2", "e1", "e2", "e3", "s1")
	fmt.Println(intCmd.Result())
	// 4 <nil>

	// 差集 返回set1中与set2不相同的元素
	strSliceCmd := rdb.SDiff(ctx, "example:set1", "example:set2")
	fmt.Println(strSliceCmd.Result())
	// [s2 s3] <nil>

	// 交集
	strSliceCmd = rdb.SInter(ctx, "example:set1", "example:set2")
	fmt.Println(strSliceCmd.Result())
	// [s1] <nil>

	// 并集
	strSliceCmd = rdb.SUnion(ctx, "example:set1", "example:set2")
	fmt.Println(strSliceCmd.Result())
	// [s3 s1 e3 e1 s2 e2] <nil>

	// 将交集存储至storeInter集合中, 并返回交集个数
	intCmd = rdb.SInterStore(ctx, "storeInter", "example:set1", "example:set2")
	fmt.Println(rdb.SMembers(ctx, "storeInter").Result())
	// [s1] <nil>
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 将差集存储至storeDiff集合中, 并返回差集个数
	intCmd = rdb.SDiffStore(ctx, "storeDiff", "example:set1", "example:set2")
	fmt.Println(rdb.SMembers(ctx, "storeDiff").Result())
	// [s2 s3] <nil>
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 将并集存储至storeUnion集合中, 并返回并集个数
	intCmd = rdb.SUnionStore(ctx, "storeUnion", "example:set1", "example:set2")
	fmt.Println(rdb.SMembers(ctx, "storeUnion").Result())
	// [s3 s1 e3 e1 s2 e2] <nil>
	fmt.Println(intCmd.Result())
	// 6 <nil>

	// 判断是否是集合中的元素
	boolCmd := rdb.SIsMember(ctx, "example:set1", "s1")
	fmt.Println(boolCmd.Result())
	// true <nil>

	// 获取成员集合数
	intCmd = rdb.SCard(ctx, "example:set1")
	fmt.Println(intCmd.Result())
	// 3 <nil>

	// 获取全部成员
	strSliceCmd = rdb.SMembers(ctx, "example:set1")
	fmt.Println(strSliceCmd.Result())
	//[s2 s3 s1] <nil>

	// 获取1个随机成员
	strCmd := rdb.SRandMember(ctx, "example:set1")
	fmt.Println(strCmd.Result())
	// s2 <nil>

	// 获取随机N个成员
	strSliceCmd = rdb.SRandMemberN(ctx, "example:set1", 2)
	fmt.Println(strSliceCmd.Result())
	// [s2 s3] <nil>

	// 移除并返回一个随机成员
	strCmd = rdb.SPop(ctx, "example:set1")
	fmt.Println(strCmd.Result())
	// s1 <nil>

	// 移除指定成员, 返回移除成功的成员个数
	intCmd = rdb.SRem(ctx, "example:set1", "s1", "s2")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	fmt.Println(rdb.SMembers(ctx, "example:set2").Result())
	// [e3 e1 s1 e2] <nil>
	// 扫描返回值匹配e*的1个元素
	scanCmd := rdb.SScan(ctx, "example:set2", 0, "e*", 1)
	var res []string
	res, cursor, err := scanCmd.Result()
	fmt.Println(res, cursor, err)
	// [e3] 4 <nil>

	// 继续从上次游标处开始扫描
	scanCmd = rdb.SScan(ctx, "example:set2", cursor, "e*", 1)
	fmt.Println(scanCmd.Result())
	// [e1] 2 <nil>
}
