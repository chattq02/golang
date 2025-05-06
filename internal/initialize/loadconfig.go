package initialize

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"Go/global"

)

func LoadConfig() {

	// Lấy thư mục làm việc hiện tại
	workingDir, errDir := os.Getwd()
	if errDir != nil {
		panic(fmt.Errorf("error getting working directory: %w", errDir))
	}
	
	// Xây dựng đường dẫn tuyệt đối đến thư mục config
	configPath := filepath.Join(workingDir, "config")
	
	viper := viper.New()
	viper.AddConfigPath(configPath) // Sử dụng đường dẫn tuyệt đối
	viper.SetConfigName("local")    // Tên file config không có phần mở rộng
	viper.SetConfigType("yaml")     // Loại file config
	
	// đọc file config

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("error reading config file: %w", err))
	}

	// read server configuration

	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Port::", viper.GetString("security.jwt.key"))

	// config struct

	//Gắn giá trị cấu hình vào cấu trúc toàn cục
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}
}