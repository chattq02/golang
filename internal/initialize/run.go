package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"Go/global"
)

func Run() *gin.Engine {
	LoadConfig()  // đọc các cấu hình trong file config
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username)
	InitLogger() // triển khai log
	global.Logger.Info("Config log ok!!", zap.String("ok","sucess"))
	InitMysql() //=> dùng gorm
	InitMysqlC() // dùng goose và sqlc
	InitServiceInterface()
	InitKafka() 
	InitRedis()
	r := InitRouter()
	return r
	// r.Run(":8082")
}