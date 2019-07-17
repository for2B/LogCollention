package Tools

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var mylogs *logs.BeeLogger

func init(){

	beego.BeeLogger.EnableFuncCallDepth(true)
	beego.BeeLogger.SetLogFuncCallDepth(4)

	mylogs = logs.NewLogger(int64(10))
	mylogs.EnableFuncCallDepth(true)
	mylogs.SetLogFuncCallDepth(4)
	mylogs.Async() //异步
	//level := beego.AppConfig.String("logsLevel")

	err := mylogs.SetLogger(logs.AdapterConsole,
		`{"color":true}`)
	if err!=nil{
		panic("SetLogger fail"+err.Error())
	}
}

//Log 输出日志
func log(level, v interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	switch level {
	case "error":
		mylogs.Error(format, v)
	case "warning":
		mylogs.Warning(format, v)

	case "info":
		mylogs.Info(format, v)
	case "debug":
		mylogs.Debug(format, v)
	default:
		mylogs.Info(format, v)
	}
}

//设置日志输出等级
func LogError(v ...interface{}) {
	log("error", v)
}
func LogWarning(v ...interface{}) {
	log("warning", v)
}
func LogInfo(v ...interface{}) {
	log("info", v)
}
func LogDebug(v ...interface{}) {
	log("debug", v)
}
