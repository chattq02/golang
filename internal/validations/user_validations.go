package validations

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func registerUserValidations() {

	// 1. Đăng ký custom validation function
	_ = ValidatorInstance.RegisterValidation("strong_password", func(fl validator.FieldLevel) bool {
		verify_purpose := fl.Field().String()
		return len(verify_purpose) >= 8
	})

	// 2. Đăng ký custom message
	_ = ValidatorInstance.RegisterTranslation("strong_password", TransInstance,
		func(ut ut.Translator) error {
			return ut.Add("strong_password", "{0} phải có ít nhất 8 ký tự, bao gồm số và chữ hoa", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("strong_password", fe.Field())
			return t
		},
	)

	// 3. Đăng ký validation cho trường confirm_password
	_ = ValidatorInstance.RegisterValidation("confirm_password", func(fl validator.FieldLevel) bool {
		password := fl.Parent().FieldByName("Password").String()
		confirmPassword := fl.Field().String()
		return password == confirmPassword
	})

	_ = ValidatorInstance.RegisterTranslation("confirm_password", TransInstance,
		func(ut ut.Translator) error {
			return ut.Add("confirm_password", "{0} không khớp với mật khẩu", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("confirm_password", fe.Field())
			return t
		},
	)
}
