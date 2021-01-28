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
	listen,err := net.ListenUDP("udp",&net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 30000,
	})
	if err != nil{
		fmt.Printf("listen failed.err %v\n",err)
		return
	}

	defer listen.Close()

	for  {
		var buf [1024]byte
		n,addr,err := listen.ReadFromUDP(buf[:])
		if err != nil{
			fmt.Printf("read from failed,err:%v\n",err)
			return
		}

		fmt.Println("接收到的数据：",string(buf[:n]))
		_,err = listen.WriteToUDP(buf[:n],addr)
		if err != nil{
			fmt.Printf("write to %v failed,err:%v\n",addr,err)
			return
		}
	}
}
