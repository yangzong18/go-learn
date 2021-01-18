package main

import "fmt"

//map
func main() {
	// map 声名
	var a map[string]int
	fmt.Println(a==nil)
	// map 初始化
	a = make(map[string]int,8)
	fmt.Println(a)

	//添加键值对
	a["张三"] = 90
	a["李四"] = 100
	fmt.Println(a)
	fmt.Printf("a:%#v\n",a)
	fmt.Printf("type of a:%T\n", a)

	//声明后直接赋值
	b := map[string]string{
		"王武":"男",
		"郑波":"nv",
	}

	fmt.Print(b)


}
