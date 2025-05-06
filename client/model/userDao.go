package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 定义一个UserDao结构体
// 完成对User结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) *UserDao {
	UserDao := &UserDao{pool: pool}
	return UserDao
}

// 1. 根据用户id返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 通过给定id去redis查询这个用户
	res, err := redis.String(conn.Do("HGET", "user", id))
	if err != nil {
		// 错误
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	// 这里我们需要把res反序列化成User实例
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json unmarshal err:", err)
		return
	}
	return
}

// 完成登录的校验 Login
// 1. Login完成对用户的验证
// 2. 如果用户id与pwd都正确，则返回一个user实例
// 3。 如果用户的id或pwd都有误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	// 先从userDao的连接池中去除一个链接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		fmt.Println("getUserById err:", err)
		return
	}
	// 这是证明这个用户
	if user.PassPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
