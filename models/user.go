package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`               // 用户名
	Password string `json:"password"`               // 密码
	State    int    `json:"state" gorm:"default:1"` // 用户状态
}

type TransformUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UserFollowers struct {
	ID         int
	UserId     int
	FollowerId int
	CreatedAt  int64
}

func (u *User) GetByName(username string) {
	db.Where("username = ?", username).First(&u)
}

func (u *User) GetById(id int) {
	db.Preload("Followers").First(&u, id)
}

func (u *User) Store() error {
	result := db.Create(&u)
	return result.Error
}

func (uf UserFollowers) IsFollow(userId, followerId int) bool {
	db.Where("user_id = ? and follower_id = ?", userId, followerId).Find(&uf)
	return uf.ID > 0
}

func (uf UserFollowers) Follow(userId, followerId int) bool {
	uf.UserId = userId
	uf.FollowerId = followerId
	uf.CreatedAt = time.Now().Unix()
	result := db.Create(&uf)
	return result.Error == nil
}

func (uf UserFollowers) UnFollow(userId, followerId int) bool {
	result := db.Delete(uf, map[string]interface{}{"user_id": userId, "follower_id": followerId})
	return result.Error == nil
}
