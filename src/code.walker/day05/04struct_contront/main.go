package main

import "fmt"

// 构造函数 构造一个结构体实例的函数
type person struct {
	name,city string
	age int8
}

func newPerson(name,city string,age int8) *person {
	return &person{
		name:name,
		city:city,
		age:age,
	}
}
func main() {
	p1 := newPerson("小李","北京",19)
	fmt.Printf("type:%T value:%#v",p1,p1)
	fmt.Println(p1)
}
