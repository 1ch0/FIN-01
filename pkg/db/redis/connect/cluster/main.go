package cluster

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// 在standalone模式下也可以手动创建集群连接
func main() {
	ctx := context.Background()
	clusterSlots := func(ctx context.Context) ([]redis.ClusterSlot, error) {
		slots := []redis.ClusterSlot{
			// 第一个master:slave 节点
			{
				Start: 0,
				End:   8191,
				Nodes: []redis.ClusterNode{{
					Addr: ":7000", // master
				}, {
					Addr: ":8000", // slave
				}},
			},
			// 第二个master:slave 节点
			{
				Start: 8192,
				End:   16383,
				Nodes: []redis.ClusterNode{{
					Addr: ":7001", // master
				}, {
					Addr: ":8001", // slave
				}},
			},
		}
		return slots, nil
	}

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots:  clusterSlots,
		RouteRandomly: true,
	})
	rdb.Ping(ctx)
}
