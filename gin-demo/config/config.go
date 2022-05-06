package config

import (
	"gopkg.in/ini.v1"
)

type configStruct struct {
	App   appStruct   `ini:"app"`
	Redis redisStruct `ini:"redis"`
	Log   logStruct   `ini:"log"`
	DB    dbStruct    `ini:"database"`
}

//应用配置
type appStruct struct {
	AppName   string `ini:"app_name"`
	Port      string `ini:"port"`
	JWT_TOKEN string `ini:"jwt_token"`
}

var App appStruct

func LoadConfig() {
	cfg, err := ini.LooseLoad(".env", ".env.local")
	//fmt.Print("进到这里了吗？")
	//fmt.Print(cfg)
	//fmt.Print("进到这里了吗？")
	if err != nil {
		panic(err)
	}
	config := new(configStruct)
	//fmt.Print("进到这里了吗？")
	err = cfg.MapTo(config)
	if err != nil {
		panic(err)
	}
	App = config.App
	Redis = config.Redis
	Log = config.Log
	DB = config.DB
}
