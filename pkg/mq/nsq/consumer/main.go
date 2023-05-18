package main

import "github.com/nsqio/go-nsq"

// 定义一个处理函数
func handleMessage(message *nsq.Message) error {
	// 处理消息
	return nil
}

func main() {
	// 创建一个消费者
	consumer, err := nsq.NewConsumer("topic-name", "channel-name", nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	// 设置处理函数
	consumer.AddHandler(nsq.HandlerFunc(handleMessage))

	// 连接到 nsqd
	err = consumer.ConnectToNSQD("nsqd-host:4150")

	// 等待信号，退出时关闭消费者
	<-consumer.StopChan

}
