package models

type Article struct {
	ID        uint
	Title     string
	Content   string
	UserId    uint
	User      User
	CreatedAt int64
	UpdatedAt int64
}

type TransformArticle struct {
	ID        uint                   `json:"id"`
	Title     string                 `json:"title"`
	Content   string                 `json:"content"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
	Author    map[string]interface{} `json:"author"`
}

func (article *Article) Store() error {
	return db.Create(&article).Error
}

func (article *Article) GetById(id int) {
	db.Preload("User").First(&article, id)
}

func (article *Article) Update() error {
	return db.Save(&article).Error
}

func (article *Article) Destroy() error {
	return db.Delete(&article).Error
}
