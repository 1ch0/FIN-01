package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	select {
	case <-time.After(time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
