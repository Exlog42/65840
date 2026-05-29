//go:build ignore

package main

import "fmt"

func main() {
	// 变量可以写在if条件里面，防止污染外部变量空间
	if age := 20; age < 8 {
		fmt.Println("年龄<8")
	} else if age < 16 {
		fmt.Println("年龄>=8 & 年龄<16")
	} else {
		fmt.Println("年龄>=16")
	}
}
