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
		// default: localhost:6379
		Addr: "139.198.166.89:56379",
	})

	// 向stream中添加消息，结果返回消息ID
	xAddArgs := &redis.XAddArgs{
		Stream: "example:stream",
		MaxLen: 10,
		Values: map[string]interface{}{"foodId": "10001", "foodName": "麻婆豆腐"},
	}
	strCmd := rdb.XAdd(ctx, xAddArgs)
	fmt.Println(strCmd.Result())
	// 1609083771429-0 <nil>

	// 插入第二条消息
	strCmd = rdb.XAdd(ctx, &redis.XAddArgs{Stream: "example:stream", Values: []string{"foodId", "10002", "foodName", "飘香鸡翅"}})
	fmt.Println(strCmd.Result())
	// 1609083771430-0 <nil>

	// 限制stream的长度为大约10个
	intCmd := rdb.XTrimApprox(ctx, "example:stream", 10)
	fmt.Println(intCmd.Result())
	// 0 <nil>

	// 获取stream流的长度
	intCmd = rdb.XLen(ctx, "example:stream")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 获取消息列表
	xMessageSliceCmd := rdb.XRange(ctx, "example:stream", "-", "+")
	fmt.Println(xMessageSliceCmd.Result())
	// [{1609083771429-0 map[foodId:10001 foodName:麻婆豆腐]} {1609083771430-0 map[foodId:10002 foodName:飘香鸡翅]}] <nil>

	// 反向获取消息列表, 第一个为最新的消息
	xMessageSliceCmd = rdb.XRevRange(ctx, "example:stream", "+", "-")
	fmt.Println(xMessageSliceCmd.Result())
	// [{1609083771430-0 map[foodId:10002 foodName:飘香鸡翅]} {1609083771429-0 map[foodId:10001 foodName:麻婆豆腐]}] <nil>

	// 读取给定id的下一条消息
	xReadArgs := &redis.XReadArgs{
		Streams: []string{"example:stream", "1609083771429-0"},
		Count:   1,
		Block:   5 * time.Second,
	}
	xStreamSliceCmd := rdb.XRead(ctx, xReadArgs)
	fmt.Println(xStreamSliceCmd.Result())
	// [{example:stream [{1609083771430-0 map[foodId:10002 foodName:飘香鸡翅]}]}] <nil>

	// 删除消息,把两条消息都删掉
	intCmd = rdb.XDel(ctx, "example:stream", "1609083771430-0", "1609083771429-0")
	fmt.Println(intCmd.Result())
	// 2 <nil>

	// 在stream上创建消费者组eater, 从最新的消息($表示)开始消费
	statusCmd := rdb.XGroupCreate(ctx, "example:stream", "eater", "$")
	fmt.Println(statusCmd.Result())
	// OK <nil>

	// 读取消费者组eater中的未被其他消费者读取的消息>
	// 运行后会阻塞，在redis客户端命名输入 XADD "example:stream" * foodId 1003 foodName 肥宅快乐水 会获得结果
	xReadGroupArgs := &redis.XReadGroupArgs{
		Group:    "eater",                         // 消费者组
		Consumer: "eater01",                       // 消费者，用时即创建
		Streams:  []string{"example:stream", ">"}, // stream流
		Block:    0,                               // 无限等待
		NoAck:    false,                           // 需要进行确认
	}
	xStreamSliceCmd = rdb.XReadGroup(ctx, xReadGroupArgs)
	xStreamSlice, err := xStreamSliceCmd.Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xStreamSlice)
	// [{example:stream [{1609086089189-0 map[foodId:1003 foodName:肥宅快乐水]}]}]

	// 确认消息为已处理
	intCmd = rdb.XAck(ctx, "example:stream", "eater", "1609086089189-0")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 设置最后递送的id为1609086089189-0
	statusCmd = rdb.XGroupSetID(ctx, "example:stream", "eater", "1609086089189-0")
	fmt.Println(statusCmd.Result())
	// OK <nil>

	// 查看待处理消息
	// type XPending struct {
	// 	Count     int64
	// 	Lower     string
	// 	Higher    string
	// 	Consumers map[string]int64
	// }
	xPendingCmd := rdb.XPending(ctx, "example:stream", "eater")
	fmt.Println(xPendingCmd.Result())
	// &{1 1609086342551-0 1609086342551-0 map[eater01:1]} <nil>

	// 转移消息的归属权, 将超过两分钟仍未得到处理的消息转移至消费者eater02
	xClaimArgs := &redis.XClaimArgs{
		Stream:   "example:stream",
		Group:    "eater",
		Consumer: "eater02",
		MinIdle:  2 * time.Minute,
		Messages: []string{"1609086342551-0"},
	}
	xMessageSliceCmd = rdb.XClaim(ctx, xClaimArgs)
	fmt.Println(xMessageSliceCmd.Result())
	// [] <nil> // 没有满足要求的消息

	// 查看流信息
	// type XInfoStream struct {
	//   Length          int64
	//   RadixTreeKeys   int64
	//   RadixTreeNodes  int64
	//   Groups          int64
	//   LastGeneratedID string
	//   FirstEntry      XMessage
	//   LastEntry       XMessage
	// }
	xInfoStreamCmd := rdb.XInfoStream(ctx, "example:stream")
	fmt.Println(xInfoStreamCmd.Result())
	// &{3 1 2 1 1609086342551-0 {1609082364313-0 map[foodId:10001 foodName:麻婆豆腐]} {1609086342551-0 map[foodId:1003 foodName:肥宅快乐水]}} <nil>

	// 查看消费者组消息
	// type XInfoGroup struct {
	// 	Name            string
	// 	Consumers       int64
	// 	Pending         int64
	// 	LastDeliveredID string
	// }
	xInfoGroupCmd := rdb.XInfoGroups(ctx, "example:stream")
	fmt.Println(xInfoGroupCmd.Result())
	// [{eater 0 0 1609086089189-0}] <nil>

	// 删除消费者
	intCmd = rdb.XGroupDelConsumer(ctx, "example:stream", "eater", "eater01")
	fmt.Println(intCmd.Result())
	// 1 <nil>

	// 删除消费者组
	intCmd = rdb.XGroupDestroy(ctx, "example:stream", "eater")
	fmt.Println(intCmd.Result())
	// 1 <nil>
}
