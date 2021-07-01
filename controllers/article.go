package controllers

import (
	"blog-backend/models"
	"blog-backend/pkg/response"
	"blog-backend/pkg/validate"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticleForm struct {
	Title   string `form:"title" json:"string" validate:"required,min=2"`
	Content string `form:"content" json:"content" validate:"required,min=2"`
}

func ArticleStore(c *gin.Context) {
	var form ArticleForm
	if err := c.ShouldBind(&form); err != nil {
		response.Error(c, err.Error())
		return
	}

	errors := validate.Check(form)
	if errors != nil {
		response.UnValidate(c, errors)
		return
	}

	var artilce models.Article
	artilce.Title = form.Title
	artilce.Content = form.Content
	artilce.UserId = uint(c.GetInt("UserId"))
	artilce.CreatedAt = time.Now().Unix()
	if err := artilce.Store(); err != nil {
		response.Error(c, err.Error())
		return
	}

	data := make(map[string]interface{})
	data["id"] = artilce.ID
	response.Success(c, "发布成功", data)
}
