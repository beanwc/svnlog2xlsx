package logic

import (
	"errors"
	"os/exec"
	"../config"
	. "github.com/beanwc/cclog"
)

func get_svn_log_bytes(tmpPath string) (dataBytes []byte, err error) {
	dataBytes, err = call_svn_command(tmpPath)
	return
}

func call_svn_command(tmpPath string) (dataBytes []byte, err error) {
	Debug("call_svn", "call_svn_command start")
	svn_bin := config.StdConfigManager.GetConfig().Svn_bin
	cmd := exec.Command(svn_bin, "log", "-v", "--xml", tmpPath)
	Debug("call_svn", "call_svn_command Path:", cmd.Path)
	Debug("call_svn", "call_svn_command Args:", cmd.Args)
	Debug("call_svn", "call_svn_command Env:", cmd.Env)
	Debug("call_svn", "call_svn_command Dir:", cmd.Dir)
	dataBytes, output_err := cmd.CombinedOutput()
	if nil != err {
		err = errors.New("call_svn call_svn_commond error: " + output_err.Error())
		return
	}
	return
}