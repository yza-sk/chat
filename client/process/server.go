package process

import (
	"fmt"
	"os"
)

// 显示登录成功后的界面
func showMenu() {
	fmt.Printf("恭喜***成功登录")
	fmt.Printf("1. 显示在线用户列表")
	fmt.Printf("2. 发送消息")
	fmt.Printf("3. 消息列表")
	fmt.Printf("4. 退出系统")
	fmt.Printf("请选择(1-4):")
	var key int
	fmt.Scanf("%d", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你选择了退出系统...")
		os.Exit(0)
	default:
		fmt.Printf("您输入的选项不正确")
	}
}
