package main

import (
	_ "LogCollection/Routers"
	_ "LogCollection/Tools"
	"github.com/astaxie/beego"
	"go.elastic.co/apm/module/apmbeego"
)


func main(){
	beego.RunWithMiddleWares("localhost:8080", apmbeego.Middleware())
}
