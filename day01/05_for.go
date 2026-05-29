//go:build ignore

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 20 {
		i++
		fmt.Println(i)
	}

	for {
		fmt.Println(".")
	}
}
