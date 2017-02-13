package controllers

/*
import (
	"fmt"
	"github.com/astaxie/beego"
	"statistics/models"
)

// ImageController operations for searching  Statistics
type ImageController struct {
	beego.Controller
}

// URLMapping ...
func (c *ImageController) URLMapping() {
	c.Mapping("GetSex", c.GetSex)

}

// Get ...
// @Title Get Sex Image 获取性别图
// @Description get Sex Image
// @Success 200 {object} models.Statistics
// @Failure 403 :id is empty
// @router / [get]
func (c *ImageController) GetSex() {
	defer c.ServeJSON()
	sex := c.GetString("sex")
	fmt.Println(sex)
	image := models.SexMsgDatas(sex)

	c.Ctx.Output.Body(image)
}*/
