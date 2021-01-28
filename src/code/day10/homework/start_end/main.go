package main

import (
	"fmt"
	"time"
)

func main()  {
	start := time.Now()
	for i:=0;i<1000;i++ {
		//fmt.Println(i)
	}
	end := time.Now()

	fmt.Println(end.Sub(start))
}
