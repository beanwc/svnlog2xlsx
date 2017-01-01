package logic

import (
	. "github.com/beanwc/cclog"
)

func Convert(save_file string) bool {
	dataBytes, err := call_svn_command("")
	if nil != err {
		Error("logic", "convert error:", err)
		return false
	}
	Debug("logic", "convert data:", string(dataBytes))
	data, err := parse_xml(dataBytes)
	if nil != err {
		Error("logic", "convert error:", err)
		return false
	}
	Debug("logic", "convert xml:", data)
	err = generate_xlsx(data, save_file)
	if nil != err {
		Error("logic", "convert error:", err)
		return false
	}
	return true
}
