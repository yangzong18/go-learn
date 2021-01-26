package main

import "fmt"

//为任意类型添加方法
type MyInt int

func (m MyInt) sayHello() {
	fmt.Println("Hello")
}
func main() {
	var m1 MyInt
	m1 = 100
	m1.sayHello()
}
