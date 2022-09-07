package object

import "home-reward/server/helper"

var GlobalConfig = &Config{}

type Config struct {
	DataPath          string `json:"data_path"`
	TaskFileName      string `json:"task_file_name"`
	CharacterFileName string `json:"character_file_name"`
	ProductFileName   string `json:"product_file_name"`
}

func InitConfig() {
	var err error
	helper.ConfigFile, err = helper.InitFile("server/config/default.json")
	if err != nil {
		panic(err)
	}
	err = helper.ConfigFile.ToObject(GlobalConfig)
	if err != nil {
		panic(err)
	}
}
