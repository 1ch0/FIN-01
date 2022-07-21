package main

import (
	"github.com/go-redis/redis/v8"
)

// 根据传递的不同options对应返回不同的client
func main() {
	rdb1 := redis.NewUniversalClient(&redis.UniversalOptions{
		// 传入的addrs切片长度大于等于2，将返回一个集群客户端ClusterClient
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	defer rdb1.Close()

	rdb2 := redis.NewUniversalClient(&redis.UniversalOptions{
		// 传递了MasterName参数，将返回一个基于sentinel的FailoverClient
		MasterName: "master",
		Addrs:      []string{":26379"},
	})
	defer rdb2.Close()

	rdb3 := redis.NewUniversalClient(&redis.UniversalOptions{
		// addrs 切片长度为1， 将返回一个普通的单节点客户端NewClient
		Addrs: []string{":6379"},
	})
	defer rdb3.Close()
}
