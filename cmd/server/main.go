package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"Go/internal/initialize"
	_ "Go/cmd/swag/docs"

)

// @title           API documentation Ecommerce Backend SHOPDEVGO
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/chattq/Ecommerce

// @contact.name   TEAM DEV
// @contact.url    https://github.com/chattq/Ecommerce
// @contact.email  baotuyet927@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /v1/2024
// @scheme http


func main()  {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8082")
}

