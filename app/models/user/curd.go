package user

import (
	"github.com/yangliang4488/goblog/pkg/logger"
	"github.com/yangliang4488/goblog/pkg/model"
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
	return true
}

func Get(uid string) (user User, err error) {
	// todo
	return User{}, nil
}

func GetByEmail(email string) (user User, err error) {
	// todo
	return User{}, nil
}
