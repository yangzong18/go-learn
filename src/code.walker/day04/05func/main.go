package main

import (
	"fmt"
	"strings"
)

// 匿名函数
func main() {
	sayHello := func (){
		fmt.Println("匿名函数")
	}

	sayHello()


	// 立即执行
	func (){
		fmt.Println("立即执行")
	}()

	r3 := addr()
	r3() //相当于执行了内部的func

	r4 := b() // 此时是一个闭包
	r4() //相当于执行了内部的func

	r5 := c("山西农大") // 此时是一个闭包
	r5() //相当于执行了内部的func

	jpgFunc := makeSuffixFunc(".jpg")
	fmt.Println(jpgFunc("test"))

	// 进阶3
	x,y := calc(100)
	ret1 := x(200) // base = 100 + 200
	fmt.Println(ret1)
	ret2 := y(200) // base = base -200 = 300 - 200
	fmt.Println(ret2)
}


//闭包函数 进阶2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包进阶函数3
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func addr() func() {
	return func() {
		fmt.Println("hello walker")
	}
}

func b() func() {
	name := "youlike"
	return func (){
		fmt.Println("hello",name)
	}
}


func c(name string) func(){
	return func() {
		fmt.Println("hello",name)
	}
}


