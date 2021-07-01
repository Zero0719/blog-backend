package controllers

import (
	"blog-backend/models"
	"blog-backend/pkg/jwt"
	"blog-backend/pkg/response"
	"blog-backend/pkg/util"
	"blog-backend/pkg/validate"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Username string `form:"username" json:"username" validate:"required,min=6,max=20,alphanum"`
	Password string `form:"password" json:"password" validate:"required,min=6,max=20,alphanum"`
}

func UserRegister(c *gin.Context) {
	var form RegisterForm
	if err := c.ShouldBind(&form); err != nil {
		response.Error(c, err.Error())
		return
	}

	errors := validate.Check(&form)
	if errors != nil {
		response.UnValidate(c, errors)
		return
	}

	// 检查用户是否已经存在
	var user models.User
	user.GetByName(form.Username)
	if user.ID > 0 {
		response.Forbidden(c, "用户已存在")
		return
	}

	// 需要对密码进行一次加密再入库
	password := util.EncryptPassword(form.Password)

	user.Username = form.Username
	user.Password = password
	if err := user.Store(); err != nil {
		response.Error(c, err.Error())
		return
	}

	token, err := jwt.GenerateToken(int(user.ID))

	if err != nil {
		response.Error(c, "生成token失败")
		return
	}

	data := make(map[string]interface{})
	data["id"] = user.ID
	data["name"] = user.Username
	data["token"] = token

	response.Success(c, "注册成功", data)
}

type LoginForm struct {
	Username string `form:"username" json:"username" validate:"required,min=6,max=20,alphanum"`
	Password string `form:"password" json:"password" validate:"required,min=6,max=20,alphanum"`
}

func UserLogin(c *gin.Context) {
	var form LoginForm

	if err := c.ShouldBind(&form); err != nil {
		response.Error(c, err.Error())
		return
	}

	errors := validate.Check(form)
	if errors != nil {
		response.UnValidate(c, errors)
		return
	}

	var user models.User
	user.GetByName(form.Username)

	if user.ID <= 0 {
		response.UnAuthorized(c, "用户不存在")
		return
	}

	if user.Password != util.EncryptPassword(form.Password) {
		response.UnAuthorized(c, "密码不正确")
		return
	}

	token, err := jwt.GenerateToken(int(user.ID))
	if err != nil {
		response.Error(c, "生成token失败")
		return
	}

	data := make(map[string]interface{})
	data["id"] = user.ID
	data["name"] = user.Username
	data["token"] = token

	response.Success(c, "登录成功", data)
}

func UserShow(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	paramId, _ := strconv.Atoi(id)
	user.GetById(paramId)
	if user.ID <= 0 {
		response.NotFound(c, "用户不存在")
		return
	}

	data := make(map[string]interface{})

	var transformUser models.TransformUser
	transformUser.ID = user.ID
	transformUser.Username = user.Username
	data["user"] = transformUser
	response.Success(c, "请求成功", data)
}

func UserFollow(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.Error(c, err.Error())
		return
	}

	userId := c.GetInt("UserId")
	if userId <= 0 {
		response.UnAuthorized(c, "用户未登录")
		return
	}

	if id == userId {
		response.Forbidden(c, "不允许对自己操作")
		return
	}

	var uf models.UserFollowers

	if uf.IsFollow(userId, id) {
		response.Forbidden(c, "已关注该用户")
		return
	}

	if !uf.Follow(userId, id) {
		response.Error(c, "关注失败")
		return
	}

	response.Success(c, "关注成功", "")
}

func UserUnFollow(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.Error(c, err.Error())
		return
	}

	userId := c.GetInt("UserId")

	if userId == id {
		response.Forbidden(c, "不允许对自己操作")
		return
	}

	var uf models.UserFollowers

	if !uf.IsFollow(userId, id) {
		response.Forbidden(c, "未关注该用户")
		return
	}

	if !uf.UnFollow(userId, id) {
		response.Error(c, "取关失败")
		return
	}

	response.Success(c, "取关成功", "")
}
