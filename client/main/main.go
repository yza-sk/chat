package main

import (
	"fmt"
	"os"
)

var (
	userId  int
	userpwd string
)

func main() {
	// 接受用户选择
	var key int
	// 判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("——————————欢迎登录多人聊天系统——————————")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	// 根据用户的输入，显示新的信息提示
	if key == 1 {
		// 用户登录
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userpwd)
		// 先把登录的函数写道另外一个文件
		// 这里我们需要重新调用
		login(userId, userpwd)
	} else {

	}
}
