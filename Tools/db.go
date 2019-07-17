package Tools

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	err := orm.RegisterDataBase("default","mysql","root:123456@tcp(106.12.179.154:3306)/BaseDb?charset=utf8")
	if err!=nil{
		panic("open connect mysql fail"+err.Error())
	}
}