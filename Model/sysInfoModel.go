package Model

import (
	"LogCollection/Tools"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"regexp"
)

type SysInfo struct {
	SysId      string 	`orm:"pk" json:"sysid"`
	SysName    string 	`json:"sysname"`
	Maintainer string 	`json:"maintainer"`
	Telephone  string 	`json:"telephone"`
}

func init(){
	orm.RegisterModel(new(SysInfo))
	err := orm.RunSyncdb("default", false, true)
	if err!=nil{
		panic("create table fail "+ err.Error())
	}
}


func AddSysInfo(si *SysInfo)error{
	o := orm.NewOrm()

	if si.SysId == ""{
		Tools.LogInfo("SysId cannot be empty")
		return errors.New("SysId cannot be empty")
	}

	if si.SysName == ""{
		Tools.LogInfo("SysName cannot be empty")
		return errors.New("SysName cannot be empty")
	}

	if si.Maintainer == ""{
		Tools.LogInfo("maintainer cannot be empty")
		return errors.New("maintainer cannot be empty")
	}

	if si.Telephone == ""{
		Tools.LogInfo("telephone cannot be empty")
		return errors.New("telephone cannot be empty")
	}

	if !isTelephone(si.Telephone){
		Tools.LogInfo("telephone invalid")
		return errors.New("telephone invalid")
	}

	_,err:= o.Insert(si)
	if err!=nil{
		Tools.LogError("insert SysInfo fail",err)
		return err
	}
	return nil
}


func ReadSysInfo(SysId string)(SysInfo,error){

	sysInfo := SysInfo{
		SysId:SysId,
	}

	if SysId==""{
		Tools.LogInfo("SysId cannot be empty")
		return sysInfo,errors.New("SysId cannot be empty")
	}

	o := orm.NewOrm()
	err := o.Read(&sysInfo)
	if err!=nil{
		if err == orm.ErrNoRows{
			Tools.LogInfo(fmt.Sprintf("Sysid:%s does not exist",SysId))
			return sysInfo,errors.New(fmt.Sprintf("Sysid:%s does not exist",SysId))
		}else{
			Tools.LogError("Get SysInfo fail"+err.Error())
			return sysInfo,err
		}
	}

	return sysInfo,nil
}


func UpdateSysInfo(si *SysInfo)error{
	o := orm.NewOrm()

	if si.SysId == ""{
		Tools.LogInfo("SysId cannot be empty")
		return errors.New("SysId cannot be empty")
	}

	if si.Telephone!=""{
		if !isTelephone(si.Telephone){
			Tools.LogInfo("telephone invalid")
			return errors.New("telephone invalid")
		}
	}

	tempSi := SysInfo{SysId:si.SysId}
	if err := o.Read(&tempSi);err!=nil{
		if err == orm.ErrNoRows{
			Tools.LogInfo(fmt.Sprintf("Sysid :%s does not exist",si.SysId))
			return errors.New(fmt.Sprintf("Sysid:%s does not exist",si.SysId))
		}else{
			Tools.LogError("Get SysInfo fail"+err.Error())
			return err
		}
	}

	if si.SysName == ""{
		si.SysName = tempSi.SysName
	}

	if si.Telephone == ""{
		si.Telephone = tempSi.Telephone
	}

	if si.Maintainer == ""{
		si.Maintainer = tempSi.Maintainer
	}

	if _,err:=o.Update(si);err!=nil{
		Tools.LogError("Update SysInfo fail",err)
		return err
	}

	return nil
}


func DeleteSysInfo(SysId string)error{
	if SysId == ""{
		Tools.LogInfo("SysId cannot be empty")
		return errors.New("SysId cannot be empty")
	}

	o := orm.NewOrm()
	sysInfo := SysInfo{SysId:SysId}
	if _,err:=o.Delete(&sysInfo);err!=nil{
		Tools.LogError("Delete SysInfo fail",err)
		return err
	}
	return nil
}


func isTelephone(phone string)bool{
	res,err := regexp.MatchString("^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\\d{8}$",
		phone)
	if err!=nil{
		Tools.LogError("Regexp Match Phone number fail",err)
		return res
	}
	return res
}