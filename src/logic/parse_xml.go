package logic

import (
	"encoding/xml"
	"errors"
	. "github.com/beanwc/cclog"
)

type Log_entry struct {
	Revision    int     `xml:"revision,attr"`
	Author      string  `xml:"author"`
	Date        string  `xml:"date"`
	Msg         string  `xml:"msg"`
}

type Log_data struct {
	Log_entry_array []Log_entry `xml:"logentry"`
}

func parse_xml(dataBytes []byte) (svn_log *Log_data, err error){
	Debug("parse_xml", "parse_xml start")
	svn_log = new(Log_data)
	unmarshal_err := xml.Unmarshal(dataBytes, svn_log)
	if nil != unmarshal_err {
		err = errors.New("parse_xml unmarshal error: " + unmarshal_err.Error())
		return
	}
	Debug("parse_xml", "parse_xml finish")
	return
}