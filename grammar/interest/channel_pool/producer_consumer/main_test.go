package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConsumer_disposeData(t1 *testing.T) {
	// 实现一个用于处理数据的闭包，实现业务代码
	consumerHandler := func(jobs chan *Job) (b bool) {
		for jobs := range jobs {
			fmt.Println(jobs)
		}
		return
	}

	// new一个任务处理对象
	t := NewTask(consumerHandler)
	t.setConsumerPoolSize(500) // 500个协程同时消费

	// 根据自己的业务去生成数据通过AddData方法添加数据到生成channel，这里是100万条数据
	go func() {
		for i := 0; i < 1000000; i++ {
			job := new(Job)
			iStr := strconv.Itoa(i)
			job.Data = "定义任务数据格式" + iStr
			t.AddData(job)
		}
	}()

	// 消费者消费数据
	t.Consumer.disposeData(t.Production.Jobs)
}
