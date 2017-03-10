package controllers

import (
	"faceStatis/models"
	"github.com/astaxie/beego"
	"strconv"
)

// Operations about Statistics
type StatisticsController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description Pagination for Datatables
// @Param   body        body    models.AoData   true        "body for aoData: ipc=1&begin=1486051200&end=1486656000&sEcho=6&iColumns=4&sColumns=%2C%2C%2C&iDisplayStart=0&iDisplayLength=10&mDataProp_0=&sSearch_0=&bRegex_0=false&bSearchable_0=false&mDataProp_1=1&sSearch_1=&bRegex_1=false&bSearchable_1=true&mDataProp_2=2&sSearch_2=&bRegex_2=false&bSearchable_2=true&mDataProp_3=3&sSearch_3=&bRegex_3=false&bSearchable_3=true&sSearch=&bRegex=false"
// @Success 201 {int} models.Statistics
// @Failure 403 body is empty
// @router / [post]
func (c *StatisticsController) Post() {
	aColumns := []string{
		"Image",
		"Sex",
		"Age",
		"Created",
	}

	//Get Data between start and end
	maps, count, counts := models.Datatables(aColumns, c.Ctx.Input)

	data := make(map[string]interface{}, count)
	var output = make([][]interface{}, len(maps))
	for i, m := range maps {
		for _, v := range aColumns {
			if v == "Created" {
				output[i] = append(output[i], m.Time.Format("2006-01-02 03:04:05"))
			} else if v == "Age" {
				output[i] = append(output[i], m.Age)
			} else if v == "Sex" {
				if m.Gender == "male" {
					output[i] = append(output[i], "男")
				} else {
					output[i] = append(output[i], "女")
				}
			} else {
				output[i] = append(output[i], m.Url)
			}
		}
	}

	data["sEcho"], _ = strconv.Atoi(c.Ctx.Input.Query("sEcho"))
	data["iTotalRecords"] = counts
	data["iTotalDisplayRecords"] = count
	data["aaData"] = output
	c.Data["json"] = data
	c.ServeJSON()
}
