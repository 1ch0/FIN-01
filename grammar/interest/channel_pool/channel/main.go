package main

import "sync"

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func New(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *Pool) Done() {
	p.wg.Wait()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
