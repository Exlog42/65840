//go:build ignore

package main

import "fmt"

// 切片在go的底层源码里是一个结构体，有三个字段：
// 1. 指针，指向底层的真实数组
// 2. 长度，当前能通过切片看到/操作的元素个数
// 3. 容量，这个底层数组从切片开始算起，还有多少空余空间
func main() {
	array1 := [...]int{6, 7, 8, 9, 0}
	// 直接声明
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	// 从数组上切下来
	slice2 := array1[0:3]
	fmt.Println(slice2)

	// 使用make
	// 长度3, 容量5
	slice3 := make([]int, 3, 5)
	fmt.Println(slice3)

	// append操作
	slice1 = append(slice1, 20, 30)
	fmt.Println(slice1)

	// 切片的切片
	slice4 := slice1[0:4]
	fmt.Println(slice4)

}
