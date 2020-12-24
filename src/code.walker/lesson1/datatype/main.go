package main

import (
	"fmt"
	"math"
)

func main() {

	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	s1 := `第一行
第二行"这里不用转移"
第三行
`
	fmt.Println(s1)

	fmt.Println("'/Users/apple/go/go-learn/src/code.walker/lesson1/datatype'")
}
