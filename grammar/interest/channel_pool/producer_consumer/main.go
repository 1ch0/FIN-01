package main

import "sync"

type task struct {
	Production
	Consumer
}

func (t *task) setConsumerPoolSize(poolSize int) {
	t.Production.Jobs = make(chan *Job, poolSize*10)
	t.Consumer.WorkPoolSize = poolSize
}

type Job struct {
	Data string
}

func NewTask(handler func(jobs chan *Job) (b bool)) (t *task) {
	t = &task{
		Production: Production{Jobs: make(chan *Job, 100)},
		Consumer:   Consumer{WorkPoolSize: 10, Handler: handler},
	}
	return
}

type Production struct {
	Jobs chan *Job
}

func (t *Production) AddData(data *Job) {
	t.Jobs <- data
}

type Consumer struct {
	WorkPoolSize int
	Handler      func(jobs chan *Job) (b bool)
	Wg           sync.WaitGroup
}

func (c *Consumer) disposeData(data chan *Job) {
	for i := 0; i <= c.WorkPoolSize; i++ {
		c.Wg.Add(1)
		go func() {
			c.Wg.Done()
		}()
	}
	c.Wg.Wait()
}
