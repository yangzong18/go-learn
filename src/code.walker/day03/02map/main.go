package main

import "fmt"

func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	k,ok := scoreMap["里斯"]

	if ok {
		fmt.Print(k)
	}else{
		fmt.Print("没有这个简直")
	}

	//便利map
	//遍历map时的元素顺序与添加键值对的顺序无关
	fmt.Println()
	for k,v := range scoreMap {
		fmt.Printf("k:%v,v:%v",k,v)
	}

	// 只便利k值
	for k := range scoreMap {
		fmt.Println(k)
	}

	//只便利v值
	for _,v := range scoreMap {
		fmt.Println(v)
	}

	// 删除建制
	delete(scoreMap,"张三")
	fmt.Println(scoreMap)
}
