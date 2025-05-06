package process

import (
	"encoding/json"
	"example.com/chat/common/message"
	"example.com/chat/server/utils"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	// 遍历服务器端的onlineUsers map[int]*UserProcess
	// 将消息转发去除

	// 取出mes的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for id, up := range userMgr.OnlineUsers {
		// 这里，还需要过滤到自己，既不要再发给自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOlineUser(data []byte, conn net.Conn) {
	// 创建一个Transfer实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
