package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`

	} `mapstructure:"server"`
	DataBases []struct {
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("../../../config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// đọc file config

	err := viper.ReadInConfig()

	if err != nil {
        panic(fmt.Errorf("Error reading config file: %w \n", err))
    }

	// read server configuration

	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Port::", viper.GetString("security.jwt.key"))

	// config struct

	var config Config 
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}

	fmt.Println("Config Port::", config.Server.Port)

	for _, db := range config.DataBases {
		fmt.Printf("databases User %s, password: %s, host: %s", db.User, db.Password, db.Host)
	}
}