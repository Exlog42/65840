//go:build ignore

package main

import "fmt"

func main() {
	map01 := make(map[string]int, 5)
	map01["bob"] = 15
	map01["alice"] = 16
	fmt.Println(map01)

	delete(map01, "bob")
	fmt.Println(map01)

	if age, ok := map01["sub"]; ok {
		fmt.Println("sub's age is", age)
	} else {
		fmt.Println("sub is not exist")
	}

	// 遍历
	for name, age := range map01 {
		fmt.Println(name, ":", age)
	}

}
