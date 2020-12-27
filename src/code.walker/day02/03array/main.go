package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [3]int //定义一个长度为3的整形数组
	a = [3]int{1,2,3}
	var b [5]int
	b = [5]int{12,34,53,99,100}
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(b)

	var c = [3]string{"北京","上海","深圳"}
	fmt.Println(c)

	// [...] 编辑器数个数

	var d = [...]int{1,2,3,4,5,6,67}
	fmt.Printf("%T\n",d)

	// 根据索引初始化
	var f [20]int
	f = [20]int{19:1}
	fmt.Println(f[19])

	//数组的便利

	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}

	// for range
	for key,val := range a{
		fmt.Printf("key is %d,val is %v\n",key,val)
	}


}
