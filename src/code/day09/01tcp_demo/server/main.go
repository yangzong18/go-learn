package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close() // 处理完请求后关闭
	// 针对当前的连接，做数据的发送和接收
	for  {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n,err := reader.Read(buf[:])
		if err != nil{
			fmt.Printf("read from conn failed,err:%v\n",err)
			break
		}

		revc := string(buf[:n])
		fmt.Println("接收到的数据：",revc)
		conn.Write([]byte(revc))//接收到的数据返回到客户端
	}
}
// tcp server demo
func main() {
	// 1. 开启服务
	listen,err := net.Listen("tcp","127.0.0.1:20000")
	if err != nil{
		fmt.Printf("listen failed.err %v\n",err)
		return
	}

	for  {
		// 等待客户端进来建立连接
		conn,err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed.err %v\n",err)
			return
		}

		// 3.启动一个goroutine 去处理连接
		go process(conn)
	}
}
