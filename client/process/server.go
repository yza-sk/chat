package process

import (
	"encoding/json"
	"example.com/chat/client/utils"
	"example.com/chat/common/message"
	"fmt"
	"net"
	"os"
)

// 显示登录成功后的界面
func showMenu() {
	fmt.Println("恭喜***成功登录")
	fmt.Println("1. 显示在线用户列表")
	fmt.Println("2. 发送消息")
	fmt.Println("3. 消息列表")
	fmt.Println("4. 退出系统")
	fmt.Println("请选择(1-4):")
	var key int
	var content string

	SmsProcess := &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		//fmt.Println("发送消息")
		fmt.Println("请输入要发送的消息")
		fmt.Scanf("%s\n", &content)
		SmsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你选择了退出系统...")
		os.Exit(0)
	default:
		fmt.Println("您输入的选项不正确")
	}
}

func serverProcessMes(Conn net.Conn) {
	defer Conn.Close()
	// 创建一个transfer实例，不停的读取服务端发送的消息
	tf := &utils.Transfer{
		Conn: Conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务端发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err:", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 1. 有人上线广播
			var notifyUserStatusMes message.NotifyUserStatusMes
			fmt.Println(notifyUserStatusMes)
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			// 2. 将这个用户信息状态保存到客户map中
			updataUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回了未知的消息类型")
		}
		// 如果读到消息，进行下一步处理
		fmt.Println("mes:", mes)
	}
}
