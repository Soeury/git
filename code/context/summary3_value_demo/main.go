package main

import (
	"context"
	"fmt"
	"time"
)

// ctx 实现数据共享
//   这里最好时使用自定义的类型，比如可以这么写 :  type myString string

// 自定义变量类型
type Key string

// 指定唯一的值作为键
var processIDKey Key = "processID"
var testIDKey Key = "testID"

// 通过 context 取出值
func getID(ctx context.Context) (processID string, testID string) {

	processID, ok := ctx.Value(processIDKey).(string)
	if !ok {
		processID = "no processID"
	}

	testID, ok = ctx.Value(testIDKey).(string)
	if !ok {
		testID = "no testID"
	}
	return
}

func test(ctx context.Context, name string, id string) {

	defer fmt.Printf("%s goroutine exit...\n", name)

	// 创建 test 协程自身拥有的 k-v
	ctx = context.WithValue(ctx, testIDKey, id) // 沿用父ctx的kv , processID 和 testID 都不为空

	processID, testID := getID(ctx)
	for {
		fmt.Printf("%s goroutine processID:%s , testID:%s\n", name, processID, testID)
		time.Sleep(time.Second)
	}
}

func process(ctx context.Context, name string, id string) {

	defer fmt.Printf("%s goroutine exit...", name)

	// 创建一个增生的 valueCtx ,存储当前协程所拥有的 k-v
	ctx = context.WithValue(ctx, processIDKey, id) // processID 指定了 ， testID为空

	go test(ctx, "son1", "001")
	go test(ctx, "son2", "002")

	processID, testID := getID(ctx)
	for {
		fmt.Printf("%s goroutine processID:%s , testID:%s\n", name, processID, testID)
		time.Sleep(time.Second)
	}
}

func main() {

	ctx := context.Background()

	go process(ctx, "process", "8080")
	processID, testID := getID(ctx) // main 里面没有创建字段，所以processID,testID为空

	for {
		fmt.Printf("main processID:%s , testID:%s\n", processID, testID)
		time.Sleep(time.Second)
	}
}
