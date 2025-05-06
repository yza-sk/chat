package process

import (
	"encoding/binary"
	"encoding/json"
	"example.com/chat/client/utils"
	"example.com/chat/common/message"
	"fmt"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userName string, userPwd string) (err error) {
	// 1. 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	// 2.准备发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3.创建一个registerMes结构体
	var registerMes message.LoginMes
	registerMes.UserId = userId
	registerMes.UserPwd = userPwd
	registerMes.UserName = userName

	// 4. 将registerMes序列化
	data, err := json.Marshal(registerMes)
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

	// 创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
	mes, err = tf.ReadPkg() // mes就是

	if err != nil {
		fmt.Println("readPkg err:", err)
		return
	}

	// 将mes的Data部分反序列化成LoginResMes
	var RegisterResMes message.RegisterResMes
	//fmt.Println(mes.Data)
	err = json.Unmarshal([]byte(mes.Data), &RegisterResMes)
	if RegisterResMes.Code == 200 {
		fmt.Println("注册成功")
	} else if RegisterResMes.Code == 500 {
		fmt.Println(RegisterResMes.Error)
	}
	return
}

// 给关联一个用户登录的方法
// 写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
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
	tf := utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg() // mes就是

	if err != nil {
		fmt.Println("readPkg err:", err)
		return
	}

	// 将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	//fmt.Println(mes.Data)
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// 初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		//fmt.Println("登录成功")

		// 可以显示当前在线用户列表，遍历loginResMes.usersId
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UserIds {
			fmt.Println("userId:", v)

			// 完成客户端onlineUsers初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		// 这里我们还需要在客户端启动一个协程
		// 改协程保持和服务端的通讯，如果服务器有数据推送给客户端
		// 则接收并显示在客户端的终端
		go serverProcessMes(conn)
		// 1. 显示我们登陆成功的菜单..
		for {
			showMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return nil
}
