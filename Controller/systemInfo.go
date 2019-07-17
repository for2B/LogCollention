package Controller

import (
	"LogCollection/Model"
	. "LogCollection/Tools"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"go.elastic.co/apm"
)



type SystemInfoController struct {
	beego.Controller
}

//Get Method return SysInfo by SysId
func (si *SystemInfoController)Get() {
	//apm monitoring
	span, _ := apm.StartSpan(si.Ctx.Request.Context(), "SystemInfoController.Get", "controller")
	span.End()

	SysId := si.GetString("sysid")

	sysInfo,err := Model.ReadSysInfo(SysId)
	if err!=nil{
		Feedback(si.Ctx.Output,fmt.Sprintf("Get SysInfo fail by sysid:%s,err:%s",SysId,err),nil)
		return
	}

	Feedback(si.Ctx.Output,"Get SysInfo Success!",sysInfo)
	return
}

//Post Method Create New SysInfo Record
func (si *SystemInfoController)Post()  {
	//apm monitoring
	span, _ := apm.StartSpan(si.Ctx.Request.Context(), "SystemInfoController.Post", "controller")
	span.End()

	newSysInfo := Model.SysInfo{}
	err := json.Unmarshal(si.Ctx.Input.RequestBody,&newSysInfo)
	if err!=nil{
		LogError("parse post data fail:",err)
		Feedback(si.Ctx.Output,fmt.Sprintf("parse post data fail : %s",err),nil)
		return
	}

	err = Model.AddSysInfo(&newSysInfo)
	if err!=nil{
		Feedback(si.Ctx.Output,"add SysInfo fail : "+err.Error(),nil)
		return
	}

	Feedback(si.Ctx.Output,fmt.Sprintf("add SysInfo ID:%s Success",newSysInfo.SysId),nil)
}

//Put Method Update SysInfo by SysId
func (si *SystemInfoController)Put() {

	//apm monitoring
	span, _ := apm.StartSpan(si.Ctx.Request.Context(), "SystemInfoController.Put", "controller")
	span.End()

	sysInfo := Model.SysInfo{}
	err := json.Unmarshal(si.Ctx.Input.RequestBody,&sysInfo)
	if err!=nil{
		LogError("parse post data fail:",err)
		Feedback(si.Ctx.Output,fmt.Sprintf("parse post data fail: %s",err),nil)
		return
	}

	err = Model.UpdateSysInfo(&sysInfo)
	if err!=nil{
		Feedback(si.Ctx.Output,err.Error(),nil)
		return
	}

	Feedback(si.Ctx.Output,"Update SysInfo Success!",nil)
}

//Delete Method delete SysInfo Record by SysId
func (si *SystemInfoController)Delete(){
	//apm monitoring
	span, _ := apm.StartSpan(si.Ctx.Request.Context(), "SystemInfoController.Delete", "controller")
	span.End()

	SysId := si.GetString("sysid")

	err := Model.DeleteSysInfo(SysId)
	if err!=nil{
		Feedback(si.Ctx.Output,fmt.Sprintf("Delete SysInfo fail : %s",err.Error()),nil)
		return
	}

	Feedback(si.Ctx.Output,"Delete SysInfo Success",nil)
}


