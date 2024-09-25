package main

import (
	chron "github.com/robfig/cron/v3"
	"log"
)

func main() {
	config := Configure()
	manager := &Manager{}
	pollChron := chron.New()
	generalAppConfig, _ := config.GetSection("app")
	pollChron.AddFunc("@every "+generalAppConfig.Key("pollCadence").String()+"m", func() {
		manager.pollForCleanup(config)
	})
	pollChron.Start()
	select {}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}
