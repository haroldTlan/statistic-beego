package models

import (
	"github.com/astaxie/beego/logs"
	"runtime"
)

//logs
func AddLog(err interface{}, v ...interface{}) {
	if _, ok := err.(error); ok {
		pc, _, line, _ := runtime.Caller(1)
		logs.Error("[Info] ", runtime.FuncForPC(pc).Name(), line, v, err)
	} else {
		logs.Info("[Info] ", err)
	}
}
