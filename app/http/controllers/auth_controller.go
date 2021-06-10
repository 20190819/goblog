package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/yangliang4488/goblog/pkg/auth"
	"net/http"

	"github.com/yangliang4488/goblog/app/models/user"
	"github.com/yangliang4488/goblog/app/requests"
)

type AuthController struct{}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	//
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}
	// 校验
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		data, _ := json.MarshalIndent(errs, "", "  ")
		fmt.Fprint(w, string(data))
	} else {
		// 入库
		_user.Create()

		if _user.ID > 0 {
			fmt.Fprint(w, "插入成功，ID 为 ", _user.GetStringID())
		} else {
			fmt.Fprint(w, "插入异常")
		}
	}

}

func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {

	// session.Put("uid", "liangY11111")
	fmt.Fprint(w, "登录页111")
}

func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if err := auth.Attempt(email, password); err == nil {
		fmt.Fprint(w, "登录成功")
	} else {

	}
}
