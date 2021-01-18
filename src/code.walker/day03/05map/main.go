package main

import "fmt"
//值为切片类型的map
func main() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海","广州","深圳","成都")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
