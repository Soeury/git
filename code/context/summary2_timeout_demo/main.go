package main

import (
	"context"
	"fmt"
	"time"
)

// ctx 实现超时控制

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

// test3 不受开始创建的 timerCtx 的控制
func test3(ctx context.Context, name string) {

	defer fmt.Printf("%s goroutine exit...\n", name)

	// 右边的子协程开启一个新的 ctx , 协程控制test4
	ctx, cancel := context.WithCancel(ctx)

	go test4(ctx, "son4")

	for i := 0; i < 7; i++ {
		fmt.Printf("%s goroutine running...\n", name)
		time.Sleep(time.Second)
	}
	cancel()
}

// test4 只受 test3 的 cancel() 影响 ， 当 test3 调用取消函数后，test4 才会停止
func test4(ctx context.Context, name string) {

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

	// 使用原来的 ctx 开启新的协程 ， 这个协程不会受到 ddl 的影响，因为他是基于根 ctx 创建的
	// 在 son1 和 son2 退出后，最后随main一起执行
	go test3(ctx, "son3")

	ddl := time.Now().Add(time.Second * 3)
	ctx, cancel := context.WithDeadline(ctx, ddl) // 增生一个带有超时时间的 ctx , 到期自动发送取消信号
	defer cancel()

	go test(ctx, "son1")
	go test(ctx, "son2")

	for {
		select {
		case <-ctx.Done(): // 监听， ctx 到期后自动调用，退出程序
			return
		default:
			fmt.Printf("%s gorouinte running...\n", name)
			time.Sleep(time.Second)
		}
	}
}

func main() {

	ctx := context.Background()
	go process(ctx, "process")

	for {
		fmt.Println("main goroutine running...")
		time.Sleep(time.Second)
	}
}
