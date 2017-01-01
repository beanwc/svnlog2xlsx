package config

import (
	"github.com/BurntSushi/toml"
	. "github.com/beanwc/cclog"
)

type Col_info struct {
	Title       string
	Bg_color    string
	Fg_color    string
}

type Author_info struct {
	Author      string
	Name        string
}

type ConfigData struct {
	Svn_bin             string          `toml:"svn_bin"`
	Col_info_list       []Col_info      `toml:"col_info"`
	Author_info_list    []Author_info   `toml:"author_info"`
}

type ConfigManager struct {
	m_config    *ConfigData
}

func (this *ConfigManager) initialize() {
	this.m_config = new(ConfigData)
	_, err := toml.DecodeFile("./config/config.toml", this.m_config)
	if nil != err {
		Panic("ConfigManager", "Init", err)
	}
}

func (this *ConfigManager) GetConfig () *ConfigData {
	return this.m_config
}

func CreateConfigManager() *ConfigManager {
	manager := new(ConfigManager)
	manager.initialize()
	return manager
}

var StdConfigManager *ConfigManager