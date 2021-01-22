package main

//什么时候应该使用指针类型接收者
//1需要修改接收者中的值
//2接收者是拷贝代价比较大的大对象
//3保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

import "fmt"

// 方法的实例
type Person struct {
	name string
	age int8
}

// 构造函数
func NewPerson(name string,age int8) *Person {
	return &Person {
		name:name,
		age:age,
	}
}

// Dream 是为Person 类型定义方法
func (p Person) Dream()  {
	fmt.Printf("%s的梦想是学好go语言！\n",p.name)
}


// 指针接受者
func (p *Person) SetAge(newAge int8){
	p.age = newAge
}
// 值接受者
func (p Person) SetAge2(newAge int8){
	p.age = newAge
}
func main() {
	p1 := NewPerson("张三",int8(10))
	p1.Dream()
	// (*p1).Dream()

	// 调用修改年龄的方法
	fmt.Println(p1.age)

	p1.SetAge(20)
	fmt.Println(p1.age)

	// 调用修改年龄的方法
	fmt.Println(p1.age)

	p1.SetAge2(20)
	fmt.Println(p1.age)
}

