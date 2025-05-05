package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 读客户端发送的信息
	defer conn.Close()

	// 这里调用总控，创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Process2()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// 提示信息
	fmt.Println("服务端在8889监听。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// 监听成功之后等待客户端链接服务器
	for {
		fmt.Println("等待客户端链接服务器。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("net.Listen err:", err)
		}

		// 一旦链接成功，则启动一个携程和客户端保持通讯。。

		go process(conn)
	}
}
