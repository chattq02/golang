package validations

import (
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	ValidatorInstance *validator.Validate
	TransInstance     ut.Translator
	once              sync.Once
)

func InitValidator() {
	once.Do(func() {
		ValidatorInstance = validator.New()

		english := en.New()
		uni := ut.New(english, english)
		TransInstance, _ = uni.GetTranslator("en")

		_ = en_translations.RegisterDefaultTranslations(ValidatorInstance, TransInstance)

		RegisterAllCustomValidations()
	})
}

func GetValidator() *validator.Validate {
	return ValidatorInstance
}

func GetTranslator() ut.Translator {
	return TransInstance
}
