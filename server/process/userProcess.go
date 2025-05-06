package process

import (
	"encoding/json"
	"example.com/chat/common/message"
	"example.com/chat/server/model"
	"example.com/chat/server/utils"
	"fmt"
	"net"
)

type UserProcess struct {
	// 字段
	Conn   net.Conn
	UserId int
}

// 这里编写通知所有在线用户的方法
// userId要通知其他的在线用户
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	// 遍历online Users，然后一个一个发送
	for id, up := range userMgr.OnlineUsers {
		if id == userId {
			continue
		}
		// 开始通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	// 将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("marshal notifyUserStatusMes err:", err)
		return
	}
	// 将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	// 对mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal notifyUserStatusMes err:", err)
		return
	}
	fmt.Println("notifyUserStatusMes data:", string(data))
	// 发送，创建我们的Transfer实例，发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write notifyUserStatusMes err:", err)
		return
	}
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 核心代码
	// 1. 先从mes 中取出mes.Data, 并直接反序列化成registerMes
	var registerMes message.RegisterMes
	json.Unmarshal([]byte(mes.Data), &registerMes.User)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}

	// 2. 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterMesType

	// 在声明一个LoginResMes
	var registerResMes message.RegisterResMes

	// 我们需要到redis数据库去完成注册
	// 1. 使用model.MyUserDao到redis验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
		return
	} else {
		registerResMes.Code = 200
	}

	// 3. 将loginResMes序列化
	data, err := json.Marshal(resMes)
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

	// 我们需要到redis数据库去玩商城验证
	// 1. 使用model.MyUserDao 到redis去验证
	//fmt.Println("loginMes.UserId", loginMes.UserId)
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
		return
	} else {
		loginResMes.Code = 200
		this.UserId = loginMes.UserId
		// 通知其他在线用户该用户上线
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser(loginMes.UserId)
		// 将当前在线用户的id放入到loginMes.UsersId
		// 遍历userMgr.onlineusers
		for id, _ := range userMgr.OnlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user, "登录成功")
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
