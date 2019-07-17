package main

import (
	_ "LogCollection/Routers"
	_ "LogCollection/Tools"
	"github.com/astaxie/beego"
)


func main(){
	beego.Run()
}
