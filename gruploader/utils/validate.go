package utils

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	validate = validator.New()
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
}

func ValidateStruct(s any) error {
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return err
}
