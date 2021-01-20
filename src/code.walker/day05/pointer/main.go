package main

import "fmt"

// 指针
func main() {
	a := 10 // int 类型
	b := &a // *int 类型 （int指针）
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%v type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)

	// 变量b 是一个int类型的指针（*int）,它存储的是变量a的内存地址
	c := *b
	fmt.Println(c)

	d := 1
	modify(d)
	fmt.Println(d)
	modify2(&d)
	fmt.Println(d)
}

func modify(x int)  {
	x = 100
}

func modify2(y *int)  {
	*y = 100
}
