package main

import "fmt"

// 值接受者和指针接受者实现接口 的区别
type mover interface {
	move()
}
type person struct {
	name string
	age int8
}

//func (p person)move()  {
//	fmt.Printf("%s在跑。。。",p.name)
//}

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

// 指针接受者
func (p *person)move()  {
	fmt.Printf("%s在跑。。。",p.name)
}
func main() {
	var m mover
	//p1 := person{
	//	name:"小王子",
	//	age:13,
	//}
	p2 := &person{
		name:"walker",
		age:31,
	}
	//m = p1 // ?无法保存，因为p1是person类型的值
	//m.move()
	m = p2
	m.move()
}
