package main

import (
	"context"
	"fmt"
	"time"
)

// ctx 实现 协程取消

func test(ctx context.Context, name string) {

	defer fmt.Printf("%s goroutine exit...\n", name)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("%s goroutine running...\n", name)
			time.Sleep(time.Second)
		}
	}
}

func process(ctx context.Context, name string) {

	defer fmt.Printf("%s goroutine exit...\n", name)

	// 这里通过原有的ctx增生一个新的带有cancel功能的ctx
	ctx, cancel := context.WithCancel(ctx)

	go test(ctx, "son1")
	go test(ctx, "son2")

	for i := 0; i < 3; i++ {
		fmt.Printf("%s goroutine running...\n", name)
		time.Sleep(time.Second)
	}
	cancel() // 模拟调用取消函数
}

func main() {

	ctx := context.Background()
	go process(ctx, "process")

	for {
		fmt.Println("main goroutine running...")
		time.Sleep(time.Second)
	}
}
