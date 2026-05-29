//go:build ignore

package main

import "fmt"

func main() {
	var a int = 10
	b := 20
	// 未显示初始化的变量会被自动赋予零值
	var c int
	var d string
	var e bool
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
