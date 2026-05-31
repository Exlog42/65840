//go:build ignore

package main

import "fmt"

// go接口是一个方法签名的集合，它只规定要实现什么，但是绝不提供具体实现

type USB interface {
	Read() string
}

type Iphone struct {
	Id string
}

// Iphone实现了Read方法
func (f Iphone) Read() string {
	fmt.Println(f.Id)
	return f.Id
}

func copyData(usb USB) {
	usb.Read()
}

func main() {
	iphone1 := Iphone{
		Id: "123446",
	}
	copyData(iphone1)
}
