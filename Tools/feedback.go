package Tools

import "github.com/astaxie/beego/context"

type FeedBack struct {
	//FbCode int         `json:"code"`
	FbMsg  string      `json:"msg,omitempty"`
	FbData interface{} `json:"data,omitempty"`
}

func Feedback(c *context.BeegoOutput, msg string, data interface{}) {
	fb := &FeedBack{msg, data}
	c.JSON(fb, true, false)
	return
}