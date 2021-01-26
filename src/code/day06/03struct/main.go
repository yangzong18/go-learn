package main

import "fmt"

// 结构体的匿名字段
// 这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
type Person struct {
	 string
	 int
}
func main() {
	p1 := Person{
		"小王子",
		19,
	}

	fmt.Println(p1)
	fmt.Println(p1.string,p1.int)
}
