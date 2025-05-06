package process

import (
	"example.com/chat/client/model"
	"example.com/chat/common/message"
	"fmt"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

// 在客户端显示当前在线的用户
func outputOnlineUser() {
	// 遍历
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		fmt.Println("user id:", id)
	}
}

// 编写一个方法，处理返回的NotifyUserStatusMes
func updataUserStatus(notifuUserStatusMes *message.NotifyUserStatusMes) {

	// 适当优化
	user, ok := onlineUsers[notifuUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifuUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifuUserStatusMes.Status

	onlineUsers[notifuUserStatusMes.UserId] = user
	outputOnlineUser()
}
