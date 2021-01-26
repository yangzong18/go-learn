package main

import (
	"fmt"
	"strings"
)

//统计一个字符串中每个单词出现的次数
func main() {

	//1 定义一个map[string]int
	str := "how do you do"
	wordCount := make(map[string]int,10)


	//2 字符串中有哪些单词
	words := strings.Split(str," ")
	fmt.Println(words,wordCount)
	//3 统计单词
	for _,word := range words{
		_,ok := wordCount[word]
		if ok{
			wordCount[word] ++
		}else{
			wordCount[word] = 1
		}
	}

	fmt.Println(wordCount)
}
