package main

import (
	"example.com/chat/client/process"
	"fmt"
	"os"
)

var (
	userId   int
	userpwd  string
	userName string
)

func main() {
	// 接受用户选择
	var key int
	// 判断是否还继续显示菜单
	for {
		fmt.Println("——————————欢迎登录多人聊天系统——————————")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			// 用户登录
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userpwd)
			// 创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userpwd)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userpwd)
			fmt.Println("请输入用户名字(nickname):")
			fmt.Scanf("%s\n", &userName)
			// 2. 调用UserProcess, 完成注册的请求
			up := &process.UserProcess{}
			up.Register(userId, userName, userpwd)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}
