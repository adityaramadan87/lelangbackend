// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"lelangbackend/controllers"
	"lelangbackend/models"
	"strings"
)

var authFilter = func(context *context.Context) {
	header := strings.Split(context.Input.Header("Authorization"), " ")
	if len(header) != 2 {
		context.Abort(403, "Unauthorized")
	}

	if err := models.ValidateToken(header[1]); err != nil {
		context.Abort(403, err.Error())
	}

	fmt.Println(context.Input.URL())
}

func init() {
	//User router
	//beego.Router("/api/users/register", &controllers.UsersController{}, "post:Register")
	//beego.Router("/api/users/login", &controllers.UsersController{}, "post:Login")
	//beego.Router("/api/users/update", &controllers.UsersController{}, "post:Update")
	//beego.Router("/users/avatar/:avatarid", &controllers.UsersController{}, "*:GetUserAvatar")

	nameSpace := beego.NewNamespace("/api",
		beego.NSRouter("/users/login", &controllers.UsersController{}, "post:Login"),
		beego.NSRouter("/users/register", &controllers.UsersController{}, "post:Register"),

		beego.NSNamespace("/v1",
			beego.NSBefore(authFilter),
			beego.NSRouter("/lelang", &controllers.AuctionController{}, "get:Get"),
			beego.NSRouter("/lelang/add", &controllers.AuctionController{}, "post:Add"),
			beego.NSRouter("/lelang/bid", &controllers.AuctionController{}, "post:Bid"),
			beego.NSRouter("/lelang/bid/:auctionid", &controllers.AuctionController{}, "get:GetAllBidder"),
			beego.NSRouter("/auction/picture/:pictureid", &controllers.AuctionController{}, "*:GetAuctionPicture"),

			beego.NSRouter("/users/update", &controllers.UsersController{}, "post:Update"),
			beego.NSRouter("/users/avatar/:avatarid", &controllers.UsersController{}, "*:GetUserAvatar"),
		),
	)

	beego.AddNamespace(nameSpace)

	//LelangRouter
	//beego.Router("/api/lelang/add", &controllers.AuctionController{}, "post:Add")
	//beego.Router("/api/lelang", &controllers.AuctionController{}, "get:Get")
	//beego.Router("/api/lelang/bid", &controllers.AuctionController{}, "post:Bid")
	//beego.Router("/api/lelang/bid/:auctionid", &controllers.AuctionController{}, "get:GetAllBidder")
	//beego.Router("/auction/picture/:pictureid", &controllers.AuctionController{}, "*:GetAuctionPicture")
}
