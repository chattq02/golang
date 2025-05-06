package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"Go/global"
	"Go/internal/middlewares"
	"Go/internal/model"
	"Go/internal/service"
	"Go/pkg/response"
)

// manager controller Login User

type cUserLogin struct{}

var LoginController = new(cUserLogin)

// VerifyOTP  godoc
// @Summary      Verify OTP login by user
// @Description  Verify OTP login by user
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload  body  model.VerifyInput  true  "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData  "Internal Server Error"
// @Router       /user/verify_account [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	result, err := service.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		global.Logger.Error("Error verifying OTP", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrInvalidOTP, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}

// login godoc
// @Summary      User Login
// @Description  User Login
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        payload  body  model.LoginInput  true  "User Login Payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData  "Internal Server Error"
// @Router       /user/login [post]
func (c *cUserLogin) Login(ctx *gin.Context) {
	// Implement logic for login
	var params model.LoginInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	codeRs, dataRs, err := service.UserLogin().Login(ctx, &params)

	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeRs, dataRs)

}

// Register godoc
// @Summary      User Registration
// @Description  When a user registers, an OTP is sent to their email
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        payload  body  model.RegisterInput  true  "User Registration Payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData  "Internal Server Error"
// @Router       /user/register [post]
func (c *cUserLogin) Register(ctx *gin.Context) {
	var params model.RegisterInput

	// Lấy validation context từ gin context
	validationContext, exists := ctx.Get("validation_context")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Validator not found"})
		return
	}
	// lấy dữ liệu
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	vc := validationContext.(*middlewares.ValidationContext)
	// Validate input
	if err := vc.Validator.Struct(params); err != nil {
		errs := err.(validator.ValidationErrors)
		response.ErrorValidator(ctx, http.StatusBadRequest, errs, vc)
		return
	}

	codeStatus, err := service.UserLogin().Register(ctx, &params)

	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeStatus, nil)
}

// UpdatePasswordRegister godoc
// @Summary      UpdatePasswordRegister
// @Description  UpdatePasswordRegister
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload  body  model.UpdatePasswordRegisterInput  true  "Payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData  "Internal Server Error"
// @Router       /user/update_pass_register [post]
func (c *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params model.UpdatePasswordRegisterInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	result, err := service.UserLogin().UpdatePasswordRegister(ctx, params.UserToken, params.UserPassword)

	if err != nil {
		global.Logger.Error("Error UpdatePasswordRegister", zap.Error(err))
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
