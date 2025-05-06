package model

// 定义一个用户的结构体
type User struct {
	// 确定字段信息
	// 为了序列化和反序列化成功
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}
