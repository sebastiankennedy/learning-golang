package user

import (
	"goblog/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel

	// gorm：GORM 默认会将键小写化作为字段名称、去掉 column 项，默认是允许 NULL 的，估 default:NULL 也去掉
	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`
	// gorm："-" —— 设置 GORM 在读写时略过此字段
	PasswordConfirm string ` gorm:"-" valid:"password_confirm"`
}

// ComparePassword 对比密码是否匹配
func (u User) ComparePassword(password string) bool {
	return u.Password == password
}