package response

import (
	"Go/internal/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseData struct {
	Code    int         `json:"code"`    // status code
	Message string      `json:"message"` // thông báo lỗi
	Data    interface{} `json:"data"`    // dữ liệu trả ra
}

type ErrorResponseData struct {
	Code   int         `json:"code"`   // status code
	Err    string      `json:"err"`    // thông báo lỗi
	Detail interface{} `json:"detail"` // dữ liệu trả ra
}

type ErrorValidatorData struct {
	Code   int                                    `json:"code"`
	Errors validator.ValidationErrorsTranslations `json:"errors"`
}

// successResponse
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// errorResponse
func ErrorResponse(c *gin.Context, code int, message string) {
	// message == "" set msg[code]
	if message == "" {
		message = msg[code]
	}
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// errorValidator
func ErrorValidator(c *gin.Context, code int, errors validator.ValidationErrors, vc *middlewares.ValidationContext) {
	validationErrors := make(map[string]string)
	for _, fe := range errors {
		fieldName := fe.Field()
		// Nếu bạn dùng json tag, đổi như sau:
		if jsonTag := fe.StructField(); jsonTag != "" {
			fieldName = fe.Field()
		}
		validationErrors[fieldName] = fe.Translate(vc.Translator)
	}
	c.JSON(http.StatusOK, ErrorValidatorData{
		Code:   code,
		Errors: validationErrors,
	})
}
