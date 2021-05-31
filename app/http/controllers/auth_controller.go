package controllers

import (
	"fmt"
	"net/http"

	"github.com/yangliang4488/goblog/app/models/user"
)

type AuthController struct{}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	//
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 表单注册
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// 通过验证入库 跳转首页
	_user := user.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	_user.Create()
	// 表单不通过，重新显示表单

	if _user.ID > 0 {
		fmt.Fprint(w, "成功插入ID:", _user.GetStringID())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "创建用户失败")
	}

}
