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
	ctx := context.WithValue(context.Background(), "ctxKey1", ctxVal)
	go func(ctx context.Context) {
		data, ok := ctx.Value("ctxKey1").(map[string]string)
		if ok {
			ctxVal := make(map[string]string)
			for k, v := range data {
				ctxVal[k] = v
			}
			ctxVal["k3"] = "v3"
			ctx = context.WithValue(ctx, "ctxKey1", ctxVal)
			data, ok = ctx.Value("ctxKey1").(map[string]string)
			if !ok {
				fmt.Printf("sub goroutine get value from parent goruntine, value: %+v\n", nil)
			}
			fmt.Printf("sub goroutine get value from parent goruntine, value: %+v\n", data)
		}
	}(ctx)
	time.Sleep(time.Second)
}
