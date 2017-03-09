// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"faceStatis/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/chart",
			beego.NSInclude(
				&controllers.ChartsController{},
			),
		),
		beego.NSNamespace("/statistics",
			beego.NSInclude(
				&controllers.StatisticsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
