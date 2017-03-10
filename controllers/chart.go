package controllers

import (
	"faceStatis/models"
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
// @Param   body        body    models.AoData   true        "charts's data structure"
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
