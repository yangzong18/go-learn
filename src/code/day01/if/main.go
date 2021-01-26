package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家赌场开业！")
	}else if(age < 18){
		fmt.Println("warning....")
	}else{
		fmt.Println("成年了！")
	}

	if age2 := 28;age2>18{
		fmt.Println("成年了！")
	}


	//fmt.Println(age2)  age2 没有定义

}
