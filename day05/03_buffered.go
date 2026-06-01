//go:build ignore

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		for i := range 10 {
			ch <- i
		}
		close(ch) //及时关闭，否则读协程会阻塞
	}()

	// forrange会自动阻塞等待，如果ch被close了，循环自动结束
	for val := range ch {
		fmt.Println(val)
	}
}
