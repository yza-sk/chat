package process

import (
	"encoding/json"
	"example.com/chat/common/message"
	"fmt"
)

func outputGroupMes(mes *message.Message) {
	// 显示即可
	// 1. 反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 显示信息
	info := fmt.Sprintf("用户id:\t%d ::\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
