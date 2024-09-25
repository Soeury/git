package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 使用 context 之前 : 解决goroutine如何接收外部命令实现退出的功能

var wg sync.WaitGroup
var exit bool = false

// 这里没有解决如何通过外部命令实现退出
func worker1() {

	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("worker -1")
		time.Sleep(time.Second)
	}
}

// -1. 全局变量控制goroutine实现退出， 缺点:
//  1. 跨包调用不容易统一
//  2. goroutine中再次启动另一个goroutine，不好控制变量
func worker2() {

	defer wg.Done()
	for {
		fmt.Println("worker -2")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
}

// -2. channel 控制goroutine退出 ， 缺点: 跨包调用不统一，需要维护一个共同的channel
func worker3(exitchan chan struct{}) {

	defer wg.Done()
	for {
		fmt.Println("worker -3")
		time.Sleep(time.Second)
		select {
		case <-exitchan: // 这里接收到信号之后立即 return 函数，
			return
		default:
		}
	}
}

func worker4(ctx context.Context) {

	defer wg.Done()
	for {
		fmt.Println("worker -4")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // cancle() 执行后，会触发Done通道
			return
		default:
		}
	}
}

func worker5(ctx context.Context) {

	defer wg.Done()
	go worker6(ctx)

	for {
		fmt.Println("worker -5")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("worker -5 get ctx.Done()")
			return
		default:
		}
	}
}

func worker6(ctx context.Context) {

	for {
		fmt.Println("worker -6")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("worker -6 get ctx.Done()")
			return
		default:
		}
	}
}

func main() {

	wg.Add(1)
	go worker1()
	time.Sleep(time.Second * 3)
	wg.Wait()
	fmt.Println("worker -1 end...")
	fmt.Println("-------------------")

	wg.Add(1)
	go worker2()
	time.Sleep(time.Second * 3)
	exit = true
	wg.Wait()
	fmt.Println("worker -2 end...")
	fmt.Println("-------------------")

	wg.Add(1)
	var exitchan = make(chan struct{})
	go worker3(exitchan)
	time.Sleep(time.Second * 3)
	exitchan <- struct{}{} // struct{}{} 匿名结构体，没有字段，没有值，常用作通道发送的信号
	close(exitchan)
	wg.Wait()
	fmt.Println("worker -3 end...")
	fmt.Println("-------------------")

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go worker4(ctx)
	time.Sleep(time.Second * 3)
	cancel() // 用来取消context，此时会触发ctx.Done()通道
	wg.Wait()
	fmt.Println("worker -4 end...")
	fmt.Println("-------------------")

	wg.Add(1)
	ctx, cancel = context.WithCancel(context.Background())
	go worker5(ctx) // 这个 goroutine 中携带子 goroutine
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
	fmt.Println("worker -5 end...")
	fmt.Println("-------------------")

}
