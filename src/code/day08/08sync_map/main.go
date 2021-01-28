package main

import (
	"fmt"
	"sync"
)

var m = make(map[int]int)
// 安全map 类型
var m2 = sync.Map{}
var wg sync.WaitGroup

func get(key int) int {
	return m[key]
}

func set(key int,value int)  {
	m[key] = value
}
func main() {
	for i:=0;i<20;i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i,i+100)
			//set(i,i+100)
			//fmt.Printf("key:%v value:%v",i,get(i))
			val,_ := m2.Load(i)
			fmt.Printf("key:%v value:%v\n",i,val)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
