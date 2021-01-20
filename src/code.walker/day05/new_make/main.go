package main

import "fmt"
//二者都是用来做内存分配的。
//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
func main() {
	var a *int
	a = new(int)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b = make(map[string]int,4)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
