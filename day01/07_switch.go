//go:build ignore

package main

import "fmt"

func main() {
	switch day := "1"; day {
	case "1":
		fmt.Println("1")
	case "2":
		fmt.Println("2")
	default:
		fmt.Println("default")
	}
}
