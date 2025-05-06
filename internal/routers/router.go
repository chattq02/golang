package routers

// import (
// 	"fmt"
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 

// 	"github.com/gin-gonic/gin"

// 	"Go/internal/controller"
// 	"Go/internal/middlewares"

// )

// func AA() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before ==> AA")
// 		c.Next()
// 		fmt.Println("Alter ==> AA")
//     }
// }
// func BB() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before ==> BB")
// 		c.Next()
// 		fmt.Println("Alter ==> BB")
//     }
// }
// func CC() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before ==> CC")
// 		c.Next()
// 		fmt.Println("Alter ==> CC")
//     }
// }


// func NewRouter() *gin.Engine {
// 	r := gin.Default()

// 	// use the middleware
// 	r.Use(middlewares.AuthenMiddleware(),BB(),CC())

// 	v1 := r.Group("/v1/2024")
// 	{
// 		v1.GET("/hello", controller.NewPongController().GetUserByID)
// 		v1.POST("/user", controller.NewUserController().GetUserByID)
// 	}

// 	return r
// }
