package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	fmt.Println(m) //切片 v 存储的是指针
	for k, v := range m {
		fmt.Printf("v:%#v\n",v)
		fmt.Println(k, "=>", v.name)
	}
}
