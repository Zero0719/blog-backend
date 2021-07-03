package controllers

import (
	"blog-backend/config"
	"blog-backend/models"
	"blog-backend/pkg/response"
	"blog-backend/pkg/validate"
	"time"

	"github.com/Zero0719/go-function/convert"
	"github.com/gin-gonic/gin"
)

type ConmentForm struct {
	Conent string `form:"content" json:"content" validate:"required,min=2"`
}

func CommentStore(c *gin.Context) {
	var form ConmentForm
	if err := c.ShouldBind(&form); err != nil {
		response.Error(c, err.Error())
		return
	}

	errors := validate.Check(form)

	if errors != nil {
		response.UnValidate(c, errors)
		return
	}

	articleId := c.Param("article_id")
	var article models.Article
	article.GetById(convert.StringToInt(articleId))

	if article.ID <= 0 {
		response.NotFound(c, "文章不存在")
		return
	}

	userId := c.GetInt("UserId")

	var comment models.Comment
	comment.Content = form.Conent
	comment.UserId = uint(userId)
	comment.ArticleId = article.ID
	comment.CreatedAt = time.Now().Unix()

	if err := comment.Store(); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, "success", "")
}

type CommentQuery struct {
	Page   int `form:"page"`
	Size   int `form:"size"`
	UserId int `form:"user_id"`
}

func CommentIndex(c *gin.Context) {
	articleId := convert.StringToInt(c.Param("article_id"))
	var query CommentQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Error(c, err.Error())
		return
	}

	if query.Page <= 0 {
		query.Page = 1
	}

	if query.Size <= 0 {
		query.Size = config.Conf.PageLimit
	}

	where := make(map[string]interface{})

	if articleId > 0 {
		where["article_id"] = articleId
	}

	if query.UserId > 0 {
		where["user_id"] = query.UserId
	}

	var comment models.Comment
	comments := comment.GetList(query.Page, query.Size, where)
	count := comment.Count(where)

	var transformComment models.TransformComment
	var transformComments []models.TransformComment
	userMap := make(map[string]interface{})
	for _, item := range comments {
		transformComment.ID = item.ID
		transformComment.Content = item.Content
		userMap["id"] = item.User.ID
		userMap["username"] = item.User.Username
		transformComment.User = userMap
		transformComments = append(transformComments, transformComment)
	}

	data := make(map[string]interface{})
	data["list"] = transformComments
	data["count"] = count
	response.Success(c, "success", data)
}

func CommentDestroy(c *gin.Context) {
	id := convert.StringToInt(c.Param("id"))

	var comment models.Comment
	comment.GetById(id)
	if comment.ID <= 0 {
		response.NotFound(c, "评论不存在")
		return
	}

	userId := c.GetInt("UserId")
	if comment.UserId != uint(userId) {
		response.Forbidden(c, "权限不足")
		return
	}

	if err := comment.Destroy(); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, "success", "")

}
