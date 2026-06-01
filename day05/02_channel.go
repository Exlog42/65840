//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		res := <-ch // 阻塞等待
		fmt.Println("收到数据：", res)
	}()

	time.Sleep(3 * time.Second)
	ch <- 42
	time.Sleep(1 * time.Second)
}
