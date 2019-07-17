package Model

import (
	"github.com/astaxie/beego/orm"
	"testing"
)

func init(){
	err := orm.RegisterDataBase("test","mysql","root:123456@tcp(127.0.0.1:30900)/BaseDb?charset=utf8")
	if err!=nil{
		panic("open connect mysql fail"+err.Error())
	}
}
//代码覆盖率
func TestAddSysInfo(t *testing.T) {
	sysInfo := SysInfo{
		SysId:"test1",
		SysName:"Sys1",
		Maintainer:"chl",
		Telephone:"15602215961",
	}

	err := AddSysInfo(&sysInfo)
	if err!=nil{
		t.Error("create data fail:",err)
	}

	getSysInfo ,err := ReadSysInfo(sysInfo.SysId)
	if err!=nil{
		t.Error("get data fail:",err)
	}

	if getSysInfo.SysName != sysInfo.SysName{
		t.Error("get data fail: the sysname unequal")
	}


	updateSysInfo := SysInfo{
		SysId:sysInfo.SysId,
		Telephone:"123123",
	}

	err = UpdateSysInfo(&updateSysInfo)
	if err==nil{
		t.Error("update invalid telephone should be err")
	}

	newPhone := "15602217067"
	updateSysInfo.Telephone = newPhone

	err = UpdateSysInfo(&updateSysInfo)
	if err!=nil{
		t.Error("update fail:err:",err)
	}

	getNewSysInfo ,err := ReadSysInfo(sysInfo.SysId)
	if err!=nil{
		t.Error("get data fail:",err)
	}

	if getNewSysInfo.SysName == ""{
		t.Error("update fail,order fields are assigned null")
	}

	if getNewSysInfo.Telephone!=newPhone{
		t.Error("update data fail,new telephone should be ",newPhone)
	}

	err = DeleteSysInfo(sysInfo.SysId)
	if err!=nil{
		t.Error("delete data fail:",err)
	}

}
