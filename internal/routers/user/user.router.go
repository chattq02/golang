package user

import (
	"github.com/gin-gonic/gin"

	"Go/internal/controller/account"
	"Go/internal/middlewares"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// this is non-dependency
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHandlerNonDependency := controller.NewUserController(us)
	// userController, _ := wire.InitUserRouterHandler()


	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", account.LoginController.Register)
		userRouterPublic.POST("/verify_account", account.LoginController.VerifyOTP)
		userRouterPublic.POST("/login", account.LoginController.Login)
		userRouterPublic.POST("/update_pass_register", account.LoginController.UpdatePasswordRegister)
	}

	// private router

	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthenMiddleware())
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_infor")
	}

}