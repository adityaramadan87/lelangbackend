package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["lelangbackend/controllers:AuctionController"] = append(beego.GlobalControllerRouter["lelangbackend/controllers:AuctionController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/api/lelang",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
