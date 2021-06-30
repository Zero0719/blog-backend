package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`               // 用户名
	Password string `json:"password"`               // 密码
	State    int    `json:"state" gorm:"default:1"` // 用户状态
}

func (u *User) GetByName(username string) {
	db.Where("username = ?", username).First(&u)
}

func (u *User) Store() error {
	result := db.Create(&u)
	return result.Error
}
