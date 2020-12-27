package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("PHP\n" + "JAVA\n" + "GO\n" + "Python")
	var name = "walkersdlkjlksdjf"

	//字符串长度
	fmt.Println(len(name))

	//字符串链接
	s1 := "Baidu"
	s2 := "Google"
	fmt.Println(s1 + s2)
	s3 := fmt.Sprintf("%s---%s", s1, s2)

	fmt.Println(s3)

	//用---分割字符串
	s4 := strings.Split(s3, "---")
	fmt.Println(s4)

	// 判断字符串是否包含另一个字符
	s5 := strings.Contains(s3, "Google")
	fmt.Println(s5)

	// 判断字符串是否含有https
	s6 := strings.HasPrefix("https://www.baidu.com", "https")

	fmt.Println(s6)

	s7 := strings.HasSuffix("https://www.jd.com", "cn")
	fmt.Println(s7)

	a1 := []string{"Python", "Php", "Java", "Go"}

	fmt.Println(strings.Join(a1, " "))
}
