package process

import (
	"encoding/json"
	"example.com/chat/common/message"
	"example.com/chat/server/utils"
	"fmt"
	"net"
)

type UserProcess struct {
	// 字段
	Conn net.Conn
}

// 编写一个函数serverProcessLogin函数， 专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心代码
	// 1. 先从mes 中取出mes.Data, 并直接反序列化成LoginMes
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}

	// 2. 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 在声明一个LoginResMes
	var loginResMes message.LoginResMes

	// 如果用户id = 100 ,密码=123456， 认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200
	} else {
		// 不合法
		loginResMes.Code = 500 // 表示用户不存在
		loginResMes.Error = "该用户不存在，请注册在使用"
	}

	// 3. 将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	// 4. 将data复赋值给resMes
	resMes.Data = string(data)

	// 5. 对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	// 6. 发送data，我们将其封装到writepkg函数
	// 因为使用分层模式， 我们先创建一个Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
