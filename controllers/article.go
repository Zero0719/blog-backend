package controllers

import (
	"blog-backend/models"
	"blog-backend/pkg/response"
	"blog-backend/pkg/validate"
	"strconv"
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

func ArticleUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.Error(c, err.Error())
		return
	}

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

	var article models.Article
	article.GetById(id)

	if article.ID <= 0 {
		response.NotFound(c, "文章未找到")
		return
	}

	article.Title = form.Title
	article.Content = form.Content
	article.UpdatedAt = time.Now().Unix()

	if err := article.Update(); err != nil {
		response.Error(c, err.Error())
		return
	}

	data := make(map[string]interface{})
	data["id"] = article.ID
	response.Success(c, "修改成功", data)
}

func ArticleDestroy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.Error(c, err.Error())
		return
	}

	var article models.Article
	article.GetById(id)
	if article.ID <= 0 {
		response.NotFound(c, "文章不存在")
		return
	}

	userId := c.GetInt("UserId")

	if int(article.UserId) != userId {
		response.Forbidden(c, "权限不足")
		return
	}

	if err := article.Destroy(); err != nil {
		response.Error(c, "删除失败")
		return
	}

	response.Success(c, "删除成功", "")
}
