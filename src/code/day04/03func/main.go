package main

import "fmt"

// 函数进阶 变量作用域

var num = 100

func testGlobal()  {
	num = 19
	// 现在函数里面找，再在外边找
	fmt.Println("全局变量",num)
}
func main() {
	testGlobal()
}
