package main

import (
	"fmt"
	"time"
)

// work_pool 模式，控制goroutine的数量，防止goroutine泄漏和暴涨。

func work(id int,jobs<- chan int,results chan<- int)  {
	for job := range jobs{
		fmt.Printf("work:%d start job:%d\n",id,job)
		results<- job*2
		time.Sleep(time.Microsecond*500)
		fmt.Printf("work:%d end job:%d\n",id,job)
	}

}
func main() {
	jobs := make(chan int,100)
	results := make(chan int,200)

	// 开启3个goroutine
	for j := 0;j<3;j++ {
		go work(j,jobs,results)
	}
	for i :=0;i<5;i++ {
		jobs<- i
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		ret := <-results
		fmt.Println(ret)
	}
}
