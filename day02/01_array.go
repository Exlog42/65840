//go:build ignore

package main

import "fmt"

// 长度是类型的一部分
// 数组是值类型，默认深拷贝
func main() {
	array := [...]int{1, 2, 3, 4, 9, 8, 7, 6, 5}

	// 经典for循环遍历
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// range遍历
	for i, v := range array {
		fmt.Println(i, v)
	}
}
