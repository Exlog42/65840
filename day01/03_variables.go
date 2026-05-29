//go:build ignore

package main

import "fmt"

func main() {
	// go没有隐式类型转换
	var a int = 10
	var b int8 = 20
	var c int16 = 30
	var d string = "hello"
	var e bool = false

	// 数组，不常用
	var f [3]int = [3]int{1, 2, 3}
	// 结构体
	type user struct {
		name string
		age  int
	}
	var adam user = user{
		name: "Adam",
		age:  10,
	}

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e)
	fmt.Println("f", f)
	fmt.Println("adam", adam)

}
