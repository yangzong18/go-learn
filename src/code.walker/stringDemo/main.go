package main

import "fmt"

//字符串反转

func main()  {
	s1 := "hello"
	byteArray := []byte(s1)
	s2 := ""

	for i:=len(byteArray)-1;i>=0;i--{
		// i 是4，3，2，1，0
		s2 += string(byteArray[i])
	}

	fmt.Println(s2)

	length := len(byteArray)

	for i:=0;i<length/2;i++ {
		byteArray[i],byteArray[length-1-i] = byteArray[length-1-i],byteArray[i]
	}


	//快速互换两个值   a,b = b,a


	fmt.Println(string(byteArray))
}
