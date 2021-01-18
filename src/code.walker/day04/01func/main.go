package main

import "fmt"

func main() {
	sayHello()
	sayHello2("walker")
	name := "peter"
	sayHello2(name)
	rs := initSum(10,20)
	fmt.Println(rs)
	ret1 := initSum2()
	ret2 := initSum2(10)
	ret3 := initSum2(10, 20)
	ret4 := initSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60
	ret5 := intSum3(100)
	ret6 := intSum3(100, 10)
	ret7 := intSum3(100, 10, 20)
	ret8 := intSum3(100, 10, 20, 30)
	fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160
}

// 没有参数和返回值的函数
func sayHello()  {
	fmt.Println("hello world")
}
// 一个参数
func sayHello2(name string)  {
	fmt.Println("hello world",name)
}

// 多个参数的函数

func initSum(x int,y int) int {
	return x+y
}
// 返回值第二种写法
func initSum1(x int,y int) (ret int) {
	ret = x+y
	return
}

// 可变参数
func initSum2(x...int) int  {
	fmt.Println(x)
	sum := 0
	for _,v := range x {
		sum += v
	}
	return sum
}


// 固定参数搭配可变
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}