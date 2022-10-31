package _3_singleton

import (
	"sync"
	"testing"
)

func TestGetSingleton(t *testing.T) {
	const parentInt = 1000
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(parentInt)
	singleton := [parentInt]Singleton{}

	for i := 0; i < parentInt; i++ {
		go func(index int) {
			<-start
			singleton[index] = GetInstance()
			wg.Done()
		}(i)
	}

	close(start)
	for i := 1; i < parentInt; i++ {
		if singleton[i] != singleton[i-1] {
			t.Logf("don't compare %d", i)
		}
	}
}
