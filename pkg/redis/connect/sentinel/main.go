package sentinel

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// 并发安全的连接
func main() {
	ctx := context.Background()
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		Username:         "",
		Password:         "",
		DB:               0,
		MasterName:       "master",
		SentinelAddrs:    []string{":2378"},
		SentinelPassword: "",
	})
	rdb.Ping(ctx)
}
