//go:build ignore

package main

import "fmt"

// import "errors"

// error是一个接口
func sendError() (int, error) {
	// 普通错误
	// return 1, errors.New("这是一个错误")
	return 2, fmt.Errorf("错误编号:%d", 2)
}

func recvError() {
	val, err := sendError()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}
func main() {
	recvError()
}
