package main

import "fmt"

func main()  {
	s1 := "Golang"
	s2 := 'G'
	fmt.Println(s1,s2)

	m1 := "中国"
	m2 := '中' //utf8 3个字节

	fmt.Println(m1,m2)

	s3 := "hello沙盒"

	for i:=0;i<len(s3);i++ {
		fmt.Printf("%c\n",s3[i])
	}


	// for range 循环  是按照rune 类型去判断的
	for k,v := range s3{
		fmt.Printf("%d--%c\n",k,v)
	}

	h1 := "big"
	// 强制类型转换
	byteS1 := []byte(h1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	h2 := "白萝卜"
	runeS2 := []rune(h2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))

	h3 := "四川长虹"
	runH3 := []rune(h3)
	runH3[2] = '打'
	runH3[3] = '大'
	fmt.Println(string(runH3))
}
