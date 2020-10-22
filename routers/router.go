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

	beego.Router("/api/users", &controllers.UsersController{}, "post:Register")
	beego.Router("/api/users/login", &controllers.UsersController{}, "post:Login")
}
