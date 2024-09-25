package main

import (
	properties "gopkg.in/ini.v1"
	"log"
)

func Configure() *properties.File {
	config, err := properties.Load("/Users/aaronloomis/IdeaProjects/SystemManager/app.properties")
	if err != nil {
		log.Println("Failed to load config file.")
	}
	config.Section("dirs")
	config.Section("ttlConfig")
	return config
}
