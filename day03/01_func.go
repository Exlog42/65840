//go:build ignore

package main

import "fmt"

func main() {
	var1 := 1
	var2 := "hello"
	func1(var1, var2)
}

func func1(var1 int, var2 string) (int, int) {
	fmt.Println(var1, var2)
	return var1, var1 + 1
}
