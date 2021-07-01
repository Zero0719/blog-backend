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
	return db.Create(&article).Error
}

func (article *Article) GetById(id int) {
	db.First(&article, id)
}

func (article *Article) Update() error {
	return db.Save(&article).Error
}

func (article *Article) Destroy() error {
	return db.Delete(&article).Error
}
