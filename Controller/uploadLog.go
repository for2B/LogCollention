package Controller

import (
	. "LogCollection/Tools"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UploadLogController struct {
	beego.Controller
}

var ConnPoll Pool

//func init(){
//	var err error
//	factory  := func() (net.Conn, error) { return net.Dial("tcp", "192.168.56.1:30900") }
//	ConnPoll, err = NewChannelPool(5, 10, factory)
//	if err!=nil{
//		LogError("Create Tcp conn pool fail",err)
//	}
//}

func (ul *UploadLogController)Post(){
	Msg := struct {
		SysId string `json:"sysid"`
		Msg string `json:"msg"`
	}{}

	err := json.Unmarshal(ul.Ctx.Input.RequestBody,&Msg)
	if err!=nil{
		LogError("parse post data fail",err)
		Feedback(ul.Ctx.Output,fmt.Sprintf("parse post data fail : %s",err),nil)
		return
	}
	err = outPutTCP(Msg.Msg)
	if err!=nil{
		LogError(err)
		Feedback(ul.Ctx.Output,err.Error(),nil)
	}
	Feedback(ul.Ctx.Output,"log msg success",nil)
}

func outPutTCP(str string)error{
	//打开连接:
	conn, err := ConnPoll.Get()
	if err != nil {
		return err
	}
	//并没有真正关闭，只是放回了连接池。
	defer conn.Close()
	num, err := conn.Write([]byte(str+"\n"))
	LogInfo(fmt.Sprintf("Conn Write %v byte",num))
	return err
}