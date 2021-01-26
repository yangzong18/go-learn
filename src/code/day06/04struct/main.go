package main

import "fmt"

// 嵌套结构体
type Person struct {
	Name string
	Gender string
	Age int
	Address Address  //嵌套另一个结构体
}

type Person1 struct {
	Name string
	Gender string
	Age int
	Address  //匿名嵌套另一个结构体
}
type Address struct {
	Province string
	City string
}
func main() {
	p1 := Person{
		Name:"塔尖",
		Gender: "男",
		Age: 31,
		Address: Address{
			Province: "河北",
			City:"保定",
		},
	}

	fmt.Printf("%#v\n",p1)
	fmt.Println(p1.Address.Province)

	p2 := Person1{
		Name:"塔尖",
		Gender: "男",
		Age: 31,
		Address: Address{
			Province: "河北",
			City:"保定",
		},
	}

	//访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找
	fmt.Println(p2.Province) //匿名结构体的访问
}
