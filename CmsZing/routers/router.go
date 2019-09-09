// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"backend-cms-zing/CmsZing/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/artists",
			beego.NSInclude(
				&controllers.ArtistsController{},
			),
		),

		beego.NSNamespace("/songs",
			beego.NSInclude(
				&controllers.SongsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
