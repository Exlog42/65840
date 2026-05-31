//go:build ignore

package main

import "fmt"

func func1() {
	defer func() {
		// 调用recover捕获异常
		if err := recover(); err != nil {
			fmt.Println("警报解除，找到错误原因", err)
		}
	}() // 匿名函数

	panic("紧急错误")
	fmt.Println("这里不会被执行")
}
func main() {
	func1()
	fmt.Println("这里会被执行，因为错误已经被recover")
}
