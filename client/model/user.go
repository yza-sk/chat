package model

// 定义一个用户的结构体
type User struct {
	// 确定字段信息
	// 为了序列化和反序列化成功
	userId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	PassPwd  string `json:"pass_pwd"`
}
