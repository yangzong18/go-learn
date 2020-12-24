package main

import "fmt"
//常量单个声明
const pi = 3.1514926
const e = 2.7182
//常量批量声明
const  (
	a = 100
	b = 1000
	c
	d
)
//枚举声明
const (
	n1 = iota   // 0
	n2			// 1
	n3			// 2
	n4			// 3
)

const (
	m1 = iota
	m2
	_
	m4
)

const (
	i1 = iota //0
	i2 = 100  //100
	i3		  // 2
	i4 		 // 3
)

const i5 = iota // 0

const (
	_ = iota
	KB = 1 << (10*iota)
	MB = 1 << (10*iota)
)

func main()  {
	//pi = 3.387  常量不允许修改
	fmt.Println(pi)

	fmt.Println(a,b,c,d)
	fmt.Println(n1,n2,n3,n4)
	fmt.Println(m4)
	fmt.Println(i2,i5)
}
