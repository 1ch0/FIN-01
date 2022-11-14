package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	pool := New(100)
	for i := 0; i < 1000; i++ {
		pool.Add(1)
		go func(i int) {
			resp, err := http.Get("https://www.baidu.com")
			if err != nil {
				fmt.Println(i, err)
			} else {
				defer resp.Body.Close()
				result, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(i, string(result))
			}
			pool.Done()
		}(i)
	}
	pool.Wait()
}
