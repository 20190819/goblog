package auth

import (
	"errors"
	"fmt"
	"log"

	"github.com/yangliang4488/goblog/app/models/user"
	"github.com/yangliang4488/goblog/pkg/session"
)

func GetUid() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)

	if ok && len(uid) > 0 {
		return uid
	} else {
		return ""
	}
}

func User() user.User {
	uid := GetUid()
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		} else {
			log.Fatal(err)
		}
	}
	return user.User{}
}

func Attempt(email string, password string) error {
	_user, err := user.GetByEmail(email)
	if err != nil {
		return err
	} else {
		if !_user.ComparePassword(password) {
			errStr := "账号不存在或密码错误"
			fmt.Println(errStr)
			return errors.New(errStr)
		}
		session.Put("uid", _user.GetStringID())
		return nil
	}
}

func Login(_user user.User) {
	session.Put("uid", _user.GetStringID())
}

func Logout(_user user.User) {
	session.Forget("uid")
}

func Check() bool {
	return len(GetUid()) > 0
}
