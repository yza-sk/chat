package main

import (
	"example.com/chat/server/model"
	"fmt"
	"net"
	"time"
)

// 我们在服务器启动后，就初始化一个userDao实例
// 把它做成全局的变量，在需要和redis操作时，就直接使用即可

func process(conn net.Conn) {
	// 在创建连接后设置
	//if tcpConn, ok := conn.(*net.TCPConn); ok {
	//	tcpConn.SetKeepAlive(true)
	//	tcpConn.SetKeepAlivePeriod(30 * time.Second)
	//}
	//conn.SetWriteDeadline(time.Now().Add(30 * time.Second)) // 设置写超时
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

// 编写一个函数完成对UserDao的初始化人户
func initUserDao() {
	// 这里pool本身就是一个全局变量
	// 这里需要注意一个初始化顺序问题
	// initPool，在initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	// 当服务器启动时， 我们就去初始化redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
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
