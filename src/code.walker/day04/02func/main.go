package main

import "fmt"

// defer 语句 先进后出
func main() {
	//fmt.Println("start")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println("end")

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
func f1() int {
	x := 5
	defer func() {
		fmt.Print("f1")
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		fmt.Print("f2")
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		fmt.Print("f3")
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		fmt.Print("f4")
		x++
	}(x)
	return 5
}