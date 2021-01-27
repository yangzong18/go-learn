package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func main()  {
	for i := 0; i < 100; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go say(i)
	}
	//wg.Add(1) // 计数器+1
	//go say() // 开启了一个独立的goroutine
	fmt.Println("hello main")
	wg.Wait() // 阻塞，等待其他的goroutine 完成
	//time.Sleep(time.Second)
}

func say(i int)  {
	fmt.Println("hello say",i)
	wg.Done() // 通知wg计数器-1
}