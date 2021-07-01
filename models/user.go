package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`               // 用户名
	Password string `json:"password"`               // 密码
	State    int    `json:"state" gorm:"default:1"` // 用户状态
	//Followers []*User `gorm:"many2many:user_followers"`
}

type TransformUser struct {
	ID uint `json:"id"`
	Username string `json:"username"`
}

func (u *User) GetByName(username string) {
	db.Where("username = ?", username).First(&u)
}

func (u *User) GetById(id int) {
	db.First(&u, id)
}

func (u *User) Store() error {
	result := db.Create(&u)
	return result.Error
}
