package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["statistics/controllers:SearchController"] = append(beego.GlobalControllerRouter["statistics/controllers:SearchController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["statistics/controllers:SearchController"] = append(beego.GlobalControllerRouter["statistics/controllers:SearchController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["statistics/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["statistics/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["statistics/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["statistics/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["statistics/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["statistics/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "GetIpcAll",
			Router: `/ipcs`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["statistics/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["statistics/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["statistics/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["statistics/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["statistics/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["statistics/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
