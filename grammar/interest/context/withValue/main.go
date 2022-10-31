package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxVal := make(map[string]string)
	ctxVal["k1"] = "v1"
	ctxVal["k2"] = "v2"
	ctx := context.WithValue(context.Background(), "ctxkey1", ctxVal)
	go func(ctx context.Context) {
		data, ok := ctx.Value("ctxkey1").(map[string]string)
		if ok {
			fmt.Printf("sub goruntine get value from parent goruntine, val=%+v\n", data)
		}
	}(ctx)
	time.Sleep(time.Second)
}
