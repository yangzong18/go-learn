package main

import "fmt"
//定义结构体  可以理解成类
type person struct {
	name string
	age int
	city string
}
type person1 struct {
	name,city string
	age int
}
func main() {
	var p1 person
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}
	fmt.Println(p1.name)


	// 匿名结构体
	var user struct {
		name string
		married bool
	}

	user.name = "小公主"
	user.married = false
	fmt.Printf("%#v",user)
}
