package process

import "errors"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	OnlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		OnlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUser添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.OnlineUsers[up.UserId] = up
}

// 完成对onlineUser删除
func (this *UserMgr) DeleteOnlineUser(up *UserProcess) {
	delete(this.OnlineUsers, up.UserId)
}

// 返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.OnlineUsers
}

// 根据id返回对应的值
func (this *UserMgr) GetUserById(userId int) (up *UserProcess, err error) {
	up, ok := this.OnlineUsers[userId]
	if !ok {
		err = errors.New("user not exist")
		return
	}
	return
}
