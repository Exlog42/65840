//go:build ignore

package main

import "fmt"

// 未指定变量类型的const数字可以存储无限精度
const b = 100000000000000000000000000000000000000000000000000000000000000000

func main() {
	// 常量必须在编译时就能确定值
	const a int = 10
	var c float64 = b
	fmt.Println("a", a)
	fmt.Println("c", c)
}
