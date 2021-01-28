package main

import (
	"fmt"
	"time"
)

// select多路复用

func work(id int,jobs<- chan int,results chan<- int)  {
	for job := range jobs{
		fmt.Printf("work:%d start job:%d\n",id,job)
		results<- job*2
		time.Sleep(time.Microsecond*500)
		fmt.Printf("work:%d end job:%d\n",id,job)
	}

}
func main() {
	ch1 := make(chan int,1)
	for i:=0;i<10;i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)
			case ch1<-i:
		default:
			fmt.Println("啥都不干")
			
			
		}
	}
}
