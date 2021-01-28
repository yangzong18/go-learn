package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file,err := os.OpenFile("./walker.log",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	if err != nil{
		fmt.Println("open file failed,err:",err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i:=0;i<100;i++ {
		writer.WriteString("hello go\n")
	}

	writer.Flush()

	str := "hello 沙河"
	err = ioutil.WriteFile("./walker.log", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
