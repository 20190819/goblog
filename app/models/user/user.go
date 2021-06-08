package user

import (
	"github.com/yangliang4488/goblog/app/models"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"column:name;type:varchar(255);not null;unique;" valid:"name"`
	Email    string `gorm:"column:email;type:varchar(255);" valid:"email"`
	Password string `gorm:"column:password;type:varchar(255);" valid:"password"`

	PasswordConfirm string `gorm:"" valid:"password_confirm"`
}
