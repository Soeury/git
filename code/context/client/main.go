package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 响应数据

type RespData struct {
	Resp *http.Response
	Err  error
}

var wg sync.WaitGroup

func doCall(ctx context.Context) {

	// 请求频繁使用长连接，否则使用短链接
	transport := http.Transport{
		DisableKeepAlives: true,
	}

	// 请求频繁制造一个 client
	client := http.Client{
		Transport: &transport,
	}

	// 创建一个请求的通道
	respChan := make(chan *RespData, 1)

	// 模拟发送请求
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/", nil)
	if err != nil {
		fmt.Printf("new request failed : %s\n", err)
		return
	}

	// 创建一个新的带 ctx 超时响应的请求
	req = req.WithContext(ctx)
	wg.Add(1)
	defer wg.Wait()

	//  处理客户端请求，
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v , err:%v\n", resp, err)
		rd := &RespData{
			Resp: resp,
			Err:  err,
		}

		// 将返回的响应放入到通道中
		respChan <- rd
		wg.Done()
	}()

	// select 多路复用 实现超时控制 :
	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case ret := <-respChan:
		fmt.Println("call api success")
		if ret.Err != nil {
			fmt.Printf("call api failed , err:%s\n", err)
			return
		}
		defer ret.Resp.Body.Close()
		bs, err := io.ReadAll(req.Response.Body)
		if err != nil {
			fmt.Printf("io.readall failed:%s\n", err)
			return
		}
		fmt.Printf("resp:%v\n", string(bs))
	default:
	}
}

func main() {

	dtime := time.Millisecond * 100
	ctx, cancel := context.WithTimeout(context.Background(), dtime)
	defer cancel()
	doCall(ctx)
}
