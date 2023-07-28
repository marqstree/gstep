package config

import (
	"encoding/json"
	util_log "github.com/marqstree/gstep/util/log"
	"log"
	"os"
)

type Configuration struct {
	Port string
	Db   struct {
		Database string
		Host     string
		Port     string
		User     string
		Password string
	}
}

// 全局配置
var Config = &Configuration{}

func Setup() {
	//将配置文件:config.json中的配置读取到Config
	file, err := os.Open("config.json")
	if err != nil {
		log.Printf("cannot open file config.log: %v", err)
		panic(err)
	}

	decoder := json.NewDecoder(file)
	Config = &Configuration{}
	err = decoder.Decode(Config)
	if err != nil {
		log.Printf("decode config.log failed: %v", err)
		panic(err)
	}

	log.Printf("global config:")
	util_log.PrintPretty(Config)
}
