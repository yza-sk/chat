package process

import (
	"encoding/json"
	"example.com/chat/common/message"
	"example.com/chat/server/utils"
	"fmt"
)

type SmsProcess struct {
}

func (sms *SmsProcess) SendGroupMes(content string) (err error) {
	// 1. 创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2. 创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	// 3. 序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}

	mes.Data = string(data)

	// 4. 对比mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}

	// 5. 将mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	// 6. 发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
	return
}
