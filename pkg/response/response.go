package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, &ResponseStruct{
		1,
		msg,
		data,
	})
}

func Forbidden(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, &ResponseStruct{
		0,
		msg,
		[]interface{}{},
	})
}

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, &ResponseStruct{
		0,
		msg,
		[]interface{}{},
	})
}

func UnValidate(c *gin.Context, errors map[string]string) {
	c.JSON(http.StatusForbidden, &ResponseStruct{
		0,
		"参数不正确",
		map[string]interface{}{"errors": errors},
	})
}

func UnAuthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, &ResponseStruct{
		0,
		msg,
		[]interface{}{},
	})
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, &ResponseStruct{
		0,
		msg,
		[]interface{}{},
	})
}
