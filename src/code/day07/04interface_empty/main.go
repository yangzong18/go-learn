package main

import "fmt"

// 空接口
//没有定义任何方法的接口,任何类型都实现了空接口。

//空接口类型的变量可以存储任意类型的变量
type xxx interface {

}
func main() {
	var x interface{}
	x = "hello"
	x = 100
	x = false
	fmt.Println(x)

	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	studentInfo["hobby"] = []string{"篮球","足球"}
	fmt.Println(studentInfo)

	_,ok := x.(string)
	if !ok  {
		fmt.Println("不是string")
	}else{
		fmt.Println("是字符串")
	}

	switch t := x.(type) {
	case string:
		fmt.Printf("是字符串类型，value:%v\n",t)
	case int:
		fmt.Printf("是int类型，value:%v\n",t)
	case bool:
		fmt.Printf("是bool类型，value:%v\n",t)
	default:
		fmt.Println("猜不出来\n")

	}
}
