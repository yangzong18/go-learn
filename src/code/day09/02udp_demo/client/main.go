package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// udp client demo
// 客户端
func main() {
	c, err := net.DialUDP("udp",nil, &net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer c.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for  {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = c.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [1024]byte{}
		n,addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("recv from udp failed, err:%v\n", err)
			return
		}
		fmt.Printf("read from %v,msg:%v\n：",addr,string(buf[:n]))
	}

}
