//go:build ignore

package main

import "fmt"
import "time"

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("协程执行完毕")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("主协程退出")
}
