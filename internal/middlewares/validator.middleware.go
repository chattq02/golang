package middlewares

import (
	"Go/internal/validations"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidationContext struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func ValidatiorMiddleware() gin.HandlerFunc {
	validations.InitValidator()

	return func(c *gin.Context) {
		// Gộp Validator và Translator vào một object duy nhất
		vc := &ValidationContext{
			Validator:  validations.GetValidator(),
			Translator: validations.GetTranslator(),
		}

		// Đưa object vào context
		c.Set("validation_context", vc)
		c.Next()
	}
}
