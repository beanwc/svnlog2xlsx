package main

import (
	"./config"
	"./logic"
	. "github.com/beanwc/cclog"
	"time"
)

func main() {
	Debug("main", "program launch")
	config.StdConfigManager = config.CreateConfigManager()
	//Debug("main", "config:", config.StdConfigManager.GetConfig())
	result := logic.Convert("svn_log.xlsx")
	if result {
		Debug("main", "convert success")
	}else {
		Debug("main", "convert failed")
	}
	for {
		time.Sleep(time.Second)
	}
}
