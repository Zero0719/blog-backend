package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func EncryptPassword(password string) string {
	return Md5(password + "123456")
}

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
