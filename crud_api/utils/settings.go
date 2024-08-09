package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (
	// DB info
	Db         string
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
)

// 初始化时读取ini文件中的信息
func init() {
	configPath := "internal/config/config.ini"
	if configPathEnv := os.Getenv("CONFIG_PATH"); configPathEnv != "" {
		configPath = configPathEnv
	}
	file, err := ini.Load(configPath)
	if err != nil {
		fmt.Println("Configration file loading failed")
	}
	LoadDatabase(file)
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("DataBase").Key("Db").MustString("mysql")
	DbHost = file.Section("DataBase").Key("DbHost").MustString("localhost")
	DbPort = file.Section("DataBase").Key("DbPort").MustInt(3306)
	DbUser = file.Section("DataBase").Key("DbUser").MustString("root")
	DbPassword = file.Section("DataBase").Key("DbPassword").MustString("IDoNotKnow")
	DbName = file.Section("DataBase").Key("DbName").MustString("gozero")
}
