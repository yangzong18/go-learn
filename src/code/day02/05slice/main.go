package main

import "fmt"

func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	fmt.Println(d == nil)       //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	fmt.Printf("a:%T\n",a)

	var f = []int{1,2,4}

	fmt.Printf("f:%T\n",f)

	g := [3]int{1,2,3}
	fmt.Printf("g:%T\n",g)

	// 2 直接从数组得到切片

	h := [5]int{1, 2, 3, 4, 5}
	m := h[1:3]  // m := h[low:high]  m:=h[:] 所有

	//a[2:]  // 等同于 a[2:len(a)]
	//a[:3]  // 等同于 a[0:3]
	//a[:]   // 等同于 a[0:len(a)]

	//使用内置的cap()函数求切片的容量

	fmt.Printf("m:%v len(m):%v cap(m):%v\n", m, len(m), cap(m))
}
