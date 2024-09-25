package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	dtime := time.Now().Add(time.Second * 1)
	ctx, cancel := context.WithDeadline(context.Background(), dtime)
	defer cancel()

	// 这里不要加 default
	select {
	case <-time.After(time.Second * 3):
		fmt.Println("over slept...")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
