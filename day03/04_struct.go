//go:build ignore

package main

import "fmt"

type People struct {
	Name string
	Age  int
}

func main() {
	people := People{
		Name: "Adam",
		Age:  10,
	}
	fmt.Println(people)
	func1(&people)
	fmt.Println(people)
}

func func1(var1 *People) {
	var1.Age = 20
}
