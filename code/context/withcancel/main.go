package main

import (
	"context"
	"fmt"
)

// Context 返回的 context.Context 是一个 interface 接口，(即我们定义的ctx变量)，里面有四个方法
//    -1. Dedline()     ddl 到了就算ctx没有被cancel也会直接调用 ctx.Done()
//    -2. Done()        返回 chan struct{} 类型 ， 不是用来传递数据的，是用来传递取消信号的
//    -3. Err()
//    -4. Value()

// GO语言内置了两个方法 Backgroud 和 Todo，分别作为 Context 的 backgroud 和 todo
// 最开始都是以这两个内置的上下文对象作为最顶层的 partent context
// 即两种方式是创建根context，不具备任何功能，具体实践还是要依靠context包提供的With系列函数来进行派生：

// context包定义了四个 With 函数:
//    -1. WithCancel()
//    -2. WithDeadline()
//    -3. WithTimeOut()    通常用于数据库或者网络连接的超时控制
//    -4. WithValue()

// -1. func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// -2. func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
// -3. func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// -4. func WithValue(parent Context, key, val interface{}) Context

//  Cancel() 函数的用法:
//    cancel 函数的用法就是关闭 channel(ctx.Done()) 来取消信息，并且递归地取消所有的子goroutine
//    效果:  -1.  cancel() 之后，将取消信号传递给所有子goroutine
//           -2.  goroutine接收到消息就是 select 中的 <-ctx.Done() 执行

// 有四个结构体类型实现了 Context 接口，
//    -1. emptyCtx
//    -2. cancelCtx
//    -3. timerCtx
//    -4. valueCtx

//  1. emptyCtx 通常作为默认的顶级上下文使用，创建方式由 context.Background() 或者 context.TODO()
//      其他的 context 都是在这个基础上创建的，

//  2. cancelCtx 是 context 包中的一个结构体类型，用于传播取消信号和管理取消操作
//      通过  withCancel 函数和 withCancelCause 函数可以创建出 cancelCtx
//      这两个函数中调用的 propagateCancel 函数用来将当前创建的 cancelCtx 与父 context 关联起来，实现取消信号传播
//      区别就是返回的 cancel 函数在 withCancelCause 调用时可以传入 cause , 而 withCancel 默认无参(nil)

// WithCancel 演示
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	num := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- num:
				num++
			}
		}
	}()
	return dst
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for num := range gen(ctx) {
		fmt.Println(num)
		if num == 5 {
			break
		}
	}
}
