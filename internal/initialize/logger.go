package initialize

import (
	"Go/global"
	"Go/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}