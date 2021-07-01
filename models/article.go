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

func (article *Article) GetById(id int) {
	db.First(&article, id)
}

func (article *Article) Update() error {
	result := db.Save(&article)
	return result.Error
}
