package main

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

)

func main() {
	// // 2.
	// logger := zap.NewExample()
	// logger.Info("Hello")

	// //DEvelopment
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// //production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	//3.

	encoder := getEncoderLog()
	sync := getWriterSync()

	core := zapcore.NewCore(encoder,sync,zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line",1))
	logger.Error("Error log", zap.Int("line",2))

}

// formatlog
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder // biến đổi thời gian
	// ts -> time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder // tên file

	return zapcore.NewConsoleEncoder(encodeConfig)

}

//

func getWriterSync() zapcore.WriteSyncer {

	// Lấy đường dẫn thư mục chứa tệp hiện tại (main.log.go)
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Đường dẫn thư mục log cùng cấp với cmd
	logDir := filepath.Join(filepath.Dir(filepath.Dir(currentDir)), "log")

	// Tạo thư mục log nếu chưa tồn tại
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	logFile := filepath.Join(logDir, "log.txt")
	
	file,_ := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)

	syncFile := zapcore.AddSync(file)

	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
