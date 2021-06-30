package util

import (
	"blog-backend/config"
	"crypto/md5"
	"fmt"
	"io"
)

func EncryptPassword(password string) string {
	return Md5(password + config.Conf.Key)
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
