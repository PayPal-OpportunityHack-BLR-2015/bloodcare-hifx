package app

import (
	"fmt"
)

type Envs struct {
	Mode          string `env:"MODE;mandatory"`
	Port          int    `env:"PORT;mandatory;8000"`
	LogPath       string `env:"LOG_FILE;mandatory"`
	MysqlUser     string `env:"MYSQL_USER;mandatory"`
	MysqlPassword string `env:"MYSQL_PASSWORD;mandatory"`
	MysqlHost     string `env:"MYSQL_HOST;mandatory"`
	MysqlPort     string `env:"MYSQL_PORT;mandatory"`
	MysqlDbName   string `env:"MYSQL_DBNAME;mandatory"`
	RedisUrl      string `env:"REDIS_URL;mandatory"`
	AssetsUrl     string `env:"ASSETS_URL;mandatory"`
	UploadPath    string `env:"UPLOAD_PATH;mandatory"`
}

func (e Envs) Show() {
	fmt.Println("Mode : ", e.Mode)
	fmt.Println("Port : ", e.Port)
	fmt.Println("LogPath : ", e.LogPath)
	fmt.Println("MysqlUser : ", e.MysqlUser)
	fmt.Println("MysqlPassword : ", e.MysqlPassword)
	fmt.Println("MysqlHost : ", e.MysqlHost)
	fmt.Println("MysqlPort : ", e.MysqlPort)
	fmt.Println("MysqlDbName : ", e.MysqlDbName)
	fmt.Println("RedisUrl : ", e.RedisUrl)
	fmt.Println("AssetsUrl : ", e.AssetsUrl)
	fmt.Println("UploadPath : ", e.UploadPath)
}
