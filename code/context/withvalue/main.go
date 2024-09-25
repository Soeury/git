package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {

	defer wg.Done()
	key := TraceCode("trace_code")
	traceCode, ok := ctx.Value(key).(string) // 子goroutine中获得trace_code，不使用断言拿出来的类型是any
	if !ok {
		fmt.Println("invalid trace code")
	}

LOOP:
	for {
		fmt.Printf("worker , traceCode: %s\n", traceCode)
		time.Sleep(time.Millisecond * 10) //
		select {
		case <-ctx.Done():
			break LOOP // dtime 秒后自动调用
		default:
		}
	}
	fmt.Println("worker done!")
}

func main() {

	dtime := time.Millisecond * 50
	ctx, cancel := context.WithTimeout(context.Background(), dtime)
	ctx = context.WithValue(ctx, TraceCode("trace_code"), "12345")
	defer cancel()

	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	wg.Wait()
	fmt.Println("over...")
}
