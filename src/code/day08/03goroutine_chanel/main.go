package main

/*
两个goroutine
 1. 一个发送1-100的值
 2.chanel2 接受值 计算平方和
 */
import (
	"fmt"
)

// 一个发送1-100的值 存储到chan 中
func counter(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
//chan2 接受值 计算平方和
func squarer(ch1 <-chan int,ch2 chan<- int) {
	for {
		i, ok := <-ch1 // 通道关闭后再取值ok=false
		if !ok {
			break
		}
		ch2 <- i * i
	}
	close(ch2)
}

func printer(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int,100)
	ch2 := make(chan int,200)
	go counter(ch1)
	go squarer(ch1,ch2)
	printer(ch2)
}