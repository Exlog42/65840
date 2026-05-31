//go:build ignore

package main

import "fmt"

// 一旦数据被放到接口里，它就失去了独有字段和方法，只保留了接口规定的方法

type any interface {
}

type People struct {
	Name string
	Age  int
}

func func1(f any) {
	// fmt.Println(f.Name) // 报错
	fmt.Println(f.(People).Age) // 暴力拆解，接口变量.(具体类型)
	// Comma-ok语法
	if people, ok := f.(People); ok {
		fmt.Println(people.Name)
	}
}
func main() {
	people := People{
		Name: "Adam",
		Age:  24,
	}
	func1(people)
}
