package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

// 定义用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息类型
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code    int `json:"code"` // 返回状态码 500：未注册 200：登录成功
	UserIds []int
	Error   string `json:"error"` // 错误信息
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"` // 返回400 表示用户已经占用，200 表示注册成功
	Error string `json:"error"`
}

// 配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

// 增加一个SmsMes
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体
}
