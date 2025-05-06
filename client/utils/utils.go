package utils

import (
	"encoding/binary"
	"encoding/json"
	"example.com/chat/common/message"
	"fmt"
	"io"
	"net"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	// 分析它应该有那些字段
	Conn net.Conn
	Buf  [8096]byte // 这是传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取服务端发送的数据...")
	// conn.Read 在conn没有被关闭的情况下，才会阻塞
	// 如果客户端关闭了conn则，就不会阻塞
	_, err = this.Conn.Read(this.Buf[0:4])
	if err != nil {
		if err == io.EOF {
			fmt.Println("服务端正常退出")
			return
		} else {
			fmt.Println(this.Buf[0:4])
			fmt.Println("conn.Read err:", err)
			return
		}
	}
	//fmt.Println("buf:", buf[0:4])

	// 根据buf[0:4] 转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//fmt.Println("pkgLen:", pkgLen)
	// 根据pkgLen读取消息内容
	_, err = this.Conn.Read(this.Buf[0:pkgLen])
	//fmt.Println("buf:", string(buf[0:pkgLen]))
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	// 把pkgLen反序列化成 ->message.Message
	err = json.Unmarshal(this.Buf[0:pkgLen], &mes)
	//fmt.Println(mes, "yes")
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[0:4])
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	// 发送data本身
	fmt.Println(data)
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
	fmt.Println("客户端信息返回成功")
	return
}
