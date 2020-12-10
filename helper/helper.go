package helper

import (
	"github.com/astaxie/beego"
	"time"
)

type Helper struct {
}

type Res struct {
	Rc   int         `json:"rc"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func Response(rc int, msg string, data interface{}, u beego.Controller) {
	u.Data["json"] = Res{rc, msg, data}
	u.ServeJSON()
}

func TimeNow() (t string) {
	return time.Now().Format("02/01/2006 15:04:05")
}
