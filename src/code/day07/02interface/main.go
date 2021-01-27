package main

import "fmt"

//why interface

type dog struct {

}

func (d dog)say()  {
	fmt.Println("汪汪汪～")
}
type cat struct {

}

func (c cat)say()  {
	fmt.Println("喵喵喵～")
}

type person struct {
	name string
}

func (p person)say()  {
	fmt.Println("啊啊啊～")
}

// 接口不管你是什么类型，只管你实现什么方法

// 定义一个类型，一个抽象的类，只要实现了say()这个方法的类型都可以成为sayer 类型
type sayer interface {
	say()
}

// 打的函数
func da(arg sayer)  {
	arg.say() // 不管是传进来的是什么，我都要打Ta 打Ta 就会叫 就要执行Ta的say方法
}
func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)
	p1 := person{name: "walker"}
	da(p1)
	var s sayer
	c2 := cat{}
	p2 := person{
		"peter",
	}
	s = c2
	s.say()
	fmt.Println(s)
	s = p2
	s.say()
	fmt.Println(s)
}
