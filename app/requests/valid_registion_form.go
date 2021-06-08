package requests

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thedevsaddam/govalidator"
	"github.com/yangliang4488/goblog/app/models/user"
	"github.com/yangliang4488/goblog/pkg/model"
)

func init() {
	govalidator.AddCustomRule("not_exists", func(field, rule, message string, value interface{}) error {
		columns := strings.TrimPrefix(rule, "not_exists:")
		rng := strings.Split(columns, ",")
		tableName := rng[0]
		dbfield := rng[1]
		val := value.(string)
		var count int64
		model.DB.Table(tableName).Where(dbfield+"=?", val).Count(&count)
		if count > 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已被占用", val)
		}
		return nil
	})
}

func ValidateRegistrationForm(data user.User) map[string][]string {
	// 验证规则
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"password":         []string{"required", "min:6", "max:30"},
		"password_confirm": []string{"required", "min:6", "max:30"},
	}

	messages := govalidator.MapData{
		"name":             []string{"required:用户名为必填项", "alpha_num:格式错误，只允许数字和英文", "between:用户名长度需在 3~20 之间"},
		"email":            []string{"required", "min:4", "max:30", "email"},
		"password":         []string{"required", "min:6", "max:30"},
		"password_confirm": []string{"required", "min:6", "max:30"},
	}

	// 配置
	options := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	errs := govalidator.New(options).ValidateStruct()
	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}

	return errs
}
