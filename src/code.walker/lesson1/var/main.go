package main

import "fmt"

func main() {
	var school = "山西农大"
	var name = "walker"
	fmt.Printf("%s欢迎%s", school, name)

	// _ 匿名变量
	aa, _ := foo()
	fmt.Println()
	fmt.Println(aa)

	name1 := "walker"
	var name2 string = "Regon"
	//name := "Peter"

	fmt.Println(name1)
	fmt.Println(name2)
}

func foo() (string, int) {
	return "alex", 9000
}
