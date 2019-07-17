package Routers

import (
	"LogCollection/Controller"
	"github.com/astaxie/beego"
)

func init(){
	beego.Router("/sysinfo", &Controller.SystemInfoController{})
	beego.Router("/log",&Controller.UploadLogController{})
}
