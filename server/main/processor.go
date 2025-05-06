package main

import (
	"example.com/chat/common/message"
	process2 "example.com/chat/server/process"
	"example.com/chat/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 编写一个serverprocessmess 函数
// 功能： 根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		// 创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
		if err != nil {
			fmt.Println("serverProcessLogin err:", err)
			return
		}
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		// 创建一个smsprocess实例
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) Process2() (err error) {
	// 循环等客户端发送消息
	for {
		// 这里我们将读取数据包，直接封装成一个函数readPkg()，返回Message，Err
		// 创建一个Transfer实例完成读包的任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()

		//fmt.Println("process msg:", mes)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出。。")
				return err
			} else {
				fmt.Println("readPkg err:", err)
				return err
			}
		}

		//fmt.Println(mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes err:", err)
			return err
		}
	}
}
