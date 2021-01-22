package main

import "fmt"

// 指针结构体

type person struct {
	name,city string
	age int8
}
func main() {
	var p2 = new(person)
	fmt.Printf("%T\n",p2)
	p2.name = "Peter"
	p2.city = "chengdu"
	p2.age = 18


	// 和上面的写法相等
	//(*p2).name = "Peter"
	//(*p2).city = "chengdu"
	//(*p2).age = 18
	fmt.Printf("%#v\n",p2)

	// 去结构体的地址进行实例
	p3 := &person{}
	fmt.Printf("%T\n",p3)
	p3.name = "那拆"
	p3.city = "保定"
	fmt.Printf("%#v\n",p3)


	// 1.键值对初始化
	p4 := person{
		name :"皇上",
		city:"北京",
		age:19,
	}
	fmt.Printf("%#v\n",p4)

	p5 := &person{
		name :"皇上",
		city:"北京",
		age:19,
	}
	fmt.Printf("p5:%#v\n",p5)
	//2 值列表初始化

	p6 := person{
		"hello",
		"上海"	,
		18,
	}
	fmt.Printf("p5:%#v\n",p6)
}
