package models

type Article struct {
	ID        uint
	Title     string
	Content   string
	UserId    uint
	CreatedAt int64
	UpdatedAt int64
}

func (article *Article) Store() error {
	result := db.Create(&article)
	return result.Error
}
