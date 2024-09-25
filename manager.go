package main

import (
	properties "gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
	"time"
)

type Manager struct {
}

const (
	FULL_PERMISSIONS int = 0755
)

func (m *Manager) pollForCleanup(config *properties.File) {
	log.Println("System Manager: executing cleanup jobs.")
	ttlConfig, err := config.GetSection("ttl")
	checkErr(err, " Failed to get ttl config.")
	dirConfig, err := config.GetSection("dirs")
	checkErr(err, "Failed to get directory config.")
	downloadsTtl, _ := ttlConfig.Key("downloads").Int()
	deleteOlderThan(downloadsTtl, dirConfig.Key("downloads").String())
	//Needs to run with sudo permissions
	trashTtl, _ := ttlConfig.Key("trash").Int()
	deleteOlderThan(trashTtl, dirConfig.Key("trash").String())
	log.Println("System Manager: finished executing cleanup jobs.")
}

func deleteOlderThan(ttl int, basePath string) {
	os.Chmod(basePath, os.FileMode(FULL_PERMISSIONS))
	files, err := os.ReadDir(basePath)
	checkErr(err, "Failed to read from: "+basePath)
	log.Println("Cleaning directory: " + basePath)
	for _, file := range files {
		fullPath := basePath + file.Name()
		lookback := time.Duration(-1 * ttl)
		fileInfo, err := file.Info()
		checkErr(err, "Failed to retrieve file information: "+file.Name())
		err = os.Chmod(basePath, os.FileMode(FULL_PERMISSIONS))
		if !strings.HasPrefix(fileInfo.Name(), ".") && fileInfo.ModTime().Before(time.Now().Add(lookback*24*time.Hour)) {
			err := os.RemoveAll(fullPath)
			checkErr(err, "Failed to remove file: "+file.Name())
		}

	}
}
