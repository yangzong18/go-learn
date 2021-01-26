package main

import "fmt"

func main() {
	a := [...]int{1,2}
	b := a[:]
	fmt.Printf("%p\n",&a) //去找a 这个变量对应的内存地址
	fmt.Printf("%p\n",b)  //取b 指向的内存地址

	b[0] = 100

	fmt.Println(a,b)

	// 扩容策略
	fmt.Printf("b的容量：%v\n",cap(b))
	b = append(b,3,4,5,6)
	fmt.Printf("b的容量：%v\n",cap(b))

}
