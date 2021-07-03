package models

type Comment struct {
	ID        uint
	ArticleId uint
	Article   Article
	UserId    uint
	User      User
	Content   string
	CreatedAt int64
}

type TransformComment struct {
	ID      uint        `json:"id"`
	Content string      `json:"content"`
	User    interface{} `json:"user"`
}

func (comment *Comment) Store() error {
	result := db.Create(&comment)
	return result.Error
}

func (comment Comment) GetList(page, size int, where interface{}) []Comment {
	var list []Comment
	db.Preload("User").Where(where).Limit(size).Offset((page - 1) * size).Find(&list)
	return list
}

func (comment Comment) Count(where interface{}) int {
	var count int
	db.Model(&Comment{}).Where(where).Count(&count)
	return count
}

func (comment *Comment) GetById(id int) {
	db.First(&comment, id)
}

func (comment *Comment) Destroy() error {
	return db.Delete(&comment).Error
}
