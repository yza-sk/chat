package model

import (
	"example.com/chat/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
