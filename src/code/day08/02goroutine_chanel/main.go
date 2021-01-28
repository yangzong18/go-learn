package main

import (
	"fmt"
)

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	//ch := make(chan int,1) // 有缓存区通道 异步通道
	ch := make(chan int) // 无缓存区通道  同步通道
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}