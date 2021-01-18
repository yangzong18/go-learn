package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	//for _, key := range keys {
	//	fmt.Println(key, scoreMap[key])
	//}

	var vals = make([]int, 0, 200)
	for _,v := range scoreMap {
		vals = append(vals,v)
	}

	sort.Ints(vals)

	fmt.Println(vals)

}
