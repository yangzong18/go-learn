package main

import "fmt"

// 切片是引用类型
func main() {
	a := []int{1,2,3,4}
	b := a
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)

	//var c []int
	c := make([]int,3,3)// 申请内存空间
	copy(c,a) //深拷贝
	fmt.Println(c)
}
