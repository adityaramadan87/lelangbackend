// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"lelangbackend/controllers"
)

func init() {
	//User router
	beego.Router("/api/users/register", &controllers.UsersController{}, "post:Register")
	beego.Router("/api/users/login", &controllers.UsersController{}, "post:Login")
	beego.Router("/api/users/update", &controllers.UsersController{}, "post:Update")
	beego.Router("/users/avatar/:avatarid", &controllers.UsersController{}, "*:GetUserAvatar")

	//LelangRouter
	beego.Router("/api/lelang/add", &controllers.AuctionController{}, "post:Add")
	beego.Router("/api/lelang", &controllers.AuctionController{}, "get:Get")
	beego.Router("/api/lelang/bid", &controllers.AuctionController{}, "post:Bid")
	beego.Router("/api/lelang/bid/:auctionid", &controllers.AuctionController{}, "get:GetAllBidder")
}
