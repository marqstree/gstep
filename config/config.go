package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Port       string
	DbName     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
}

// 全局配置
var Config = &Configuration{}

func Setup() {
	//将配置文件:config.json中的配置读取到Config
	file, err := os.Open("config.json")
	if err != nil {
		log.Printf("cannot open file config.json: %v", err)
		panic(err)
	}

	decoder := json.NewDecoder(file)
	Config = &Configuration{}
	err = decoder.Decode(Config)
	if err != nil {
		log.Printf("decode config.json failed: %v", err)
		panic(err)
	}

	jsonStr, _ := json.Marshal(Config)
	log.Printf("global config: %s", jsonStr)
}
