package main

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

func main() {
	// 创建一个生产者
	producer, err := nsq.NewProducer("nsqd-host:4150", nsq.NewConfig())
	if err != nil {
		fmt.Println(err)
	}
	// 发布消息
	err = producer.Publish("topic-name", []byte("message-body"))

	// 关闭生产者
	producer.Stop()
}
