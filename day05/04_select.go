//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Channel 1"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Channel 2"
	}()

	// select会一直阻塞，知道任意一个case能够执行
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(4 * time.Second):
		fmt.Println("超时")
	}
}
