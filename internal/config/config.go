package config

import (
	"cryptobot/internal/data"
	"fmt"
	"github.com/BurntSushi/toml"
)

var ResultConfig data.Config
var configFile = "config.toml"

func InitConfig() {
	var conf data.Config
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		fmt.Println("File not found\tSet up default settings")
		return
	}
	fmt.Println(conf.ApiKeys.TelegramKey)
	ResultConfig = conf
}
