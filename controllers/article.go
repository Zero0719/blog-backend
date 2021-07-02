package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"blog-backend/pkg/response"
	"blog-backend/pkg/validate"
	"strconv"
	"time"

	"github.com/Zero0719/go-function/date"
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

func ArticleShow(c *gin.Context) {
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

	var transformArticle models.TransformArticle
	transformArticle.ID = article.ID
	transformArticle.Title = article.Title
	transformArticle.Content = article.Content
	transformArticle.CreatedAt = date.Date("2006-01-02 15:04:05", article.CreatedAt)
	transformArticle.UpdatedAt = date.Date("2006-01-02 15:04:05", article.UpdatedAt)
	authorMap := make(map[string]interface{})
	authorMap["id"] = article.User.ID
	authorMap["username"] = article.User.Username
	transformArticle.Author = authorMap
	data := make(map[string]interface{})
	data["article"] = transformArticle
	response.Success(c, "获取成功", data)
}

type ArticleQuery struct {
	Page   int `form:"page"`
	Size   int `form:"size"`
	UserId int `form:"user_id"`
}

func ArticleIndex(c *gin.Context) {
	var query ArticleQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Error(c, err.Error())
		return
	}

	if query.Page == 0 {
		query.Page = 1
	}

	if query.Size == 0 {
		query.Size = config.Conf.PageLimit
	}

	where := make(map[string]interface{})

	if query.UserId > 0 {
		where["user_id"] = query.UserId
	}

	var article models.Article
	articles := article.GetList(query.Page, query.Size, where)
	count := article.Count(where)

	var transformArticle models.TransformArticle
	var transformArticles []models.TransformArticle
	authorMap := make(map[string]interface{})
	for _, item := range articles {
		transformArticle.ID = item.ID
		transformArticle.Title = item.Title
		transformArticle.Content = item.Content
		transformArticle.CreatedAt = date.Date("2006-01-02 15:03:04", item.CreatedAt)
		transformArticle.UpdatedAt = date.Date("2006-01-02 15:03:04", item.UpdatedAt)
		authorMap["id"] = item.User.ID
		authorMap["username"] = item.User.Username
		transformArticle.Author = authorMap
		transformArticles = append(transformArticles, transformArticle)
	}

	data := make(map[string]interface{})
	data["list"] = transformArticles
	data["count"] = count

	response.Success(c, "success", data)
}
