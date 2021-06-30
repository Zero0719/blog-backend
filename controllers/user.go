package controllers

import (
	"blog-backend/models"
	"blog-backend/pkg/jwt"
	"blog-backend/pkg/response"
	"blog-backend/pkg/util"
	"blog-backend/pkg/validate"

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
