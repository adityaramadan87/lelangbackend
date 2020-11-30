package helper

import (
	"github.com/astaxie/beego"
	"time"
)

type Helper struct {
}

func Response(rc int, msg interface{}, u beego.Controller) {
	u.Data["json"] = map[string]interface{}{"rc": rc, "response": msg}
}

func TimeNow() (t string) {
	return time.Now().Format("02/01/2006 15:04:05")
}
