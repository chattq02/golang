package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"Go/internal/service"
	"Go/internal/vo"
	"Go/pkg/response"
)

type UserController struct{
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest

	
	if err := c.ShouldBindJSON(&params); err != nil {
		fmt.Printf("Error in ShouldBind: %s\n", err.Error())
        response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
        return
    }
	fmt.Printf("params: %+v\n", params)
	result := uc.userService.Register(params.Email,params.Purpose)
	response.SuccessResponse(c, result, nil)
}


// // controller --> service --> repo --> models -> dbs

