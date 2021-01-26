package main

import "fmt"
func main() {

	//算术运算符
	n1 := 19
	n2 := 3
	fmt.Println(n1+n2)
	fmt.Println(n1-n2)
	fmt.Println(n1*n2)
	fmt.Println(n1/n2)
	fmt.Println(n1%n2)

	// 关系运算符
	fmt.Println(n1 == n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)
	fmt.Println(n1 != n2)


	//逻辑运算符

	a := true
	b := false
	if a&&b {
		fmt.Println("真")
	}else{
		fmt.Println("假")
	}
}
