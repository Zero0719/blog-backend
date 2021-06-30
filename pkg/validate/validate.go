package validate

import (
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var validate *validator.Validate
var uni *ut.UniversalTranslator
var trans ut.Translator

func init() {
	zh := zh.New()
	uni = ut.New(zh)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

func Check(form interface{}) map[string]string {
	errors := make(map[string]string)
	err := validate.Struct(form)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors[strings.ToLower(err.Field())] = err.Translate(trans)
		}
		return errors
	}
	return nil
}
