package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {

	defer wg.Done()
	for {
		fmt.Println("connecting db...")
		time.Sleep(time.Second * 1) // 假设数据库连接花费了 1s
		select {
		case <-ctx.Done(): // 超过 dtime 时间后自动调用
			fmt.Println("time over deadline...")
			return
		default:
		}
	}
}

func main() {

	wg.Add(1)
	dtime := time.Millisecond * 50 // 设置一个 50ms 的超时
	ctx, cancel := context.WithTimeout(context.Background(), dtime)
	go worker(ctx)
	time.Sleep(time.Second * 6) // main 函数等待 6s
	cancel()
	wg.Wait()
	fmt.Println("goroutine over...")
}
