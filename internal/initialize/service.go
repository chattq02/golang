package initialize

import (
	"Go/global"
	"Go/internal/database"
	"Go/internal/service"
	"Go/internal/service/impl"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	// user Service Interface
	service.InitUserLogin(impl.NewUserLoginImpl(queries))
	
}