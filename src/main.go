package main

import (
	"./config"
	"./logic"
	. "github.com/beanwc/cclog"
)

func main() {
	Debug("main", "program launch")
	config.StdConfigManager = config.CreateConfigManager()
	//Debug("main", "config:", config.StdConfigManager.GetConfig())
	result := logic.Convert("save.xlsx")
	if result {
		Debug("main", "convert success")
	}else {
		Debug("main", "convert failed")
	}
}
