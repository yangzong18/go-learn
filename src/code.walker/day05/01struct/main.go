package main

import "fmt"

//自定义类型
type MyInt int

// 类型别名
type NewInt = int

// MyInt 基于int 类型的自定义类型
func main() {
	var i MyInt
	fmt.Printf("type:%T value:%v\n",i,i)
	var j NewInt
	fmt.Printf("type:%T value:%v\n",j,j)
}
