//go:build ignore

package main

import "fmt"

func main() {
	array1 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array1)
	func1(&array1)
	fmt.Println(array1)

	func2(array1[:])
	fmt.Println(array1)

}

func func1(ptr *[6]int) {
	(*ptr)[0] = 10
}

func func2(slice1 []int) {
	slice1[0] = 20
}
