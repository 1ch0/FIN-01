# redis 集群
TODO(1ch0): fix bug

### Redis 集群配置
```shell

docker exec -it redis9001 redis-cli -p 9001 -a KVge8oLd2t8eYVX7EwVmmxKPCDmwMty1 --cluster create 192.168.0.100:9001 192.168.0.100:9002 192.168.0.100:9003 192.168.0.100:9004 192.168.0.100:9005 192.168.0.100:9006 --cluster-replicas 1
```