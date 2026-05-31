//go:build ignore

package main

import "fmt"

func main() {
	res := sumdiv(2, 2, 4)
	fmt.Println(res)

}

func sumdiv(var1 int, var2 ...int) (res int) {
	sum := 0
	for _, val := range var2 {
		sum += val
	}
	res = sum / var1
	return
}
