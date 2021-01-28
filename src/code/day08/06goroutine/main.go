package main

import (
	"fmt"
	"sync"
)

// 多个gotoutine 并发操作全局变量x
var x int64
var wg sync.WaitGroup
var lock sync.Mutex  // 互斥锁

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}