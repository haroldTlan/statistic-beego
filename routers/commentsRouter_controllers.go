package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["faceStatis/controllers:ChartsController"] = append(beego.GlobalControllerRouter["faceStatis/controllers:ChartsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["faceStatis/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["faceStatis/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
