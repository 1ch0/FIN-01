package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.WithValue(context.Background(), "key1", "value1")
	ctx = context.WithValue(context.Background(), "key2", "value2")
	ctx = context.WithValue(context.Background(), "key3", "value3")
	go func(ctx context.Context) {
		data, ok := ctx.Value("key1").(string)
		if ok {
			fmt.Printf("sub goroutine get vaue from parent goroutine, val=%s\n", data)
		}
	}(ctx)
	time.Sleep(time.Second)
}
