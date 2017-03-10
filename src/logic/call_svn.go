package logic

import (
	"errors"
	"os/exec"
	"../config"
	. "github.com/beanwc/cclog"
)

func get_svn_log_bytes() (dataBytes []byte, err error) {
	dataBytes, err = call_svn_command()
	return
}

func call_svn_command() (dataBytes []byte, err error) {
	Debug("call_svn", "call_svn_command start")
	svn_bin := config.StdConfigManager.GetConfig().Svn_bin
	svn_repository_path := config.StdConfigManager.GetConfig().Svn_repository_path
	cmd := exec.Command(svn_bin, "log", "-v", "--xml", svn_repository_path)
	dataBytes, output_err := cmd.CombinedOutput()
	if nil != err {
		err = errors.New("call_svn call_svn_commond error: " + output_err.Error())
		return
	}
	Debug("call_svn", "call_svn_command finish")
	return
}