package user

import (
	"fmt"
	"github.com/yangliang4488/goblog/pkg/logger"
	"github.com/yangliang4488/goblog/pkg/model"
	password2 "github.com/yangliang4488/goblog/pkg/password"
)

func (user *User) Create() (err error) {
	err = model.DB.Create(&user).Error
	if err != nil {
		logger.LogError(err)
		return err
	} else {
		return nil
	}
}

func (user *User) GetStringID() int64 {
	return user.ID
}

func (user *User) ComparePassword(password string) bool {
	fmt.Println(user.Password)
	return password2.CheckHash(password, user.Password)
}

func Get(uid string) (user User, err error) {
	// todo
	return User{}, nil
}

func GetByEmail(email string) (user User, err error) {
	err = model.DB.Where("email=?", email).First(&user).Error
	if err == nil {
		return user, nil
	}
	return User{}, nil
}
