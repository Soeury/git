package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// 调用服务端api在客户端实现超时控制

func indexHandler(w http.ResponseWriter, r *http.Request) {

	seed := time.Now().Unix()
	src := rand.NewSource(seed)
	rnd := rand.New(src)
	number := rnd.Intn(2) // 随机生成 0-2之间的数字

	if number == 0 {
		time.Sleep(time.Second * 5) // 随机制造 5s 的慢响应
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")
}

func main() {

	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
