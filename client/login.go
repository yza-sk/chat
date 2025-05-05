package main

import (
	"encoding/binary"
	"encoding/json"
	"example.com/chat/client/utils"
	"example.com/chat/common/message"
	"fmt"
	"net"
)

func login(userId int, userPwd string) (err error) {
	// 1. 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	// 2.准备发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	// 3.创建一个loginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4. 将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	// 5. 把data赋给mes.Data
	mes.Data = string(data)

	// 6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	// 延时关闭
	defer conn.Close()

	// 7. data就是我们发送的消息
	// 7.1 先把data长度发送给服务器
	// 先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	// 发送长度
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	//fmt.Println("客户端发送消息长度成功,发送长度%d", buf[0:4])
	//time.Sleep(time.Second * 3)

	//fmt.Println(data)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Read err:", err)
	}

	// ---------------------------------------------
	// 接受服务端回复

	mes, err = utils.readPkg(conn) // mes就是

	if err != nil {
		fmt.Println("readPkg err:", err)
		return
	}

	// 将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	//fmt.Println(mes.Data)
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return nil
}
