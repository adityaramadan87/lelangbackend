package helper

import (
	"github.com/astaxie/beego"
)

type Helper struct {
}

func Response(rc int, msg interface{}, u beego.Controller) {
	u.Data["json"] = map[string]interface{}{"rc": rc, "response": msg}
}
