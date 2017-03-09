package controllers

import (
	"faceStatis/models"
	_ "fmt"
	"github.com/astaxie/beego"
	"strconv"
)

// Operations about Charts
type ChartsController struct {
	beego.Controller
}

// Get ...
// @Title Get
// @Description Pagination for Charts
// @Param   body        body    models.AoData   true        "body for aoData: ipc=1&begin=1486051200&end=1486656000&sEcho=6&iColumns=4&sColumns=%2C%2C%2C&iDisplayStart=0&iDisplayLength=10&mDataProp_0=&sSearch_0=&bRegex_0=false&bSearchable_0=false&mDataProp_1=1&sSearch_1=&bRegex_1=false&bSearchable_1=true&mDataProp_2=2&sSearch_2=&bRegex_2=false&bSearchable_2=true&mDataProp_3=3&sSearch_3=&bRegex_3=false&bSearchable_3=true&sSearch=&bRegex=false"
// @Success 201 {int} models.Charts
// @Failure 403 body is empty
// @router / [get]
func (c *ChartsController) Get() {
	start, _ := strconv.Atoi(c.GetString("start"))
	end, _ := strconv.Atoi(c.GetString("end"))
	data := models.Charts(start, end)
	c.Data["json"] = data
	c.ServeJSON()
}
