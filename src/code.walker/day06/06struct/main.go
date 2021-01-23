package main

import "fmt"

// 结构体的继承
//Animal 动物
type Animal struct {
	name string

}
// Dog 结构体
type  Dog struct {
	Feet int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (a *Animal)move()  {
	fmt.Printf("%s会跑～\n",a.name)
}

func (d *Dog)wang()  {
	fmt.Printf("%s会汪汪～\n",d.name)
}
func main() {
	d1 := &Dog{
		Feet:4,
		Animal:&Animal{//注意嵌套的是结构体指针
			name:"旺财",
		},
	}
	d1.move()
	d1.wang()
}
