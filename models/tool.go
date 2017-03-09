package models

import (
	_ "fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

/*
* aColumns []string `SQL Columns to display`
* thismodel interface{} `SQL model to use`
* ctx *context.Context `Beego ctx which contains httpcontext`
* maps []orm.Params `return result in a interface map as []orm.Params`
* count int64 `return iTotalDisplayRecords`
* counts int64 `return iTotalRecords`
*
 */

type Datatable struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type Detail struct {
	Status      string
	Description []Human
}
type Human struct {
	Url    string    `json:"url"`
	Gender string    `json:"gender"`
	Age    int       `json:"age"`
	Time   time.Time `json:"time"`
}

func Datatables(aColumns []string, Input *context.BeegoInput) (human []Human, count int64, counts int64) {
	/*分页请求
	iDisplayStart  起始数目
	iDisplayLength 每页显示数量
	*/
	iDisplayStart, _ := strconv.Atoi(Input.Query("iDisplayStart"))
	iDisplayLength, _ := strconv.Atoi(Input.Query("iDisplayLength"))

	cond := orm.NewCondition()

	/*timestamp begin to end*/
	end, _ := strconv.Atoi(Input.Query("end"))
	begin, _ := strconv.Atoi(Input.Query("begin"))

	if begin > end {
		logs.Warn("Error: Invalid time. Begin must smaller than end")
	} else {
		if begin != 0 || end != 0 {
			begins := time.Unix(int64(begin), 0)
			cond = cond.And("timestamp__gte", begins.Format("2006-01-02 03:04:05"))
			ends := time.Unix(int64(end), 0)
			cond = cond.And("timestamp__lte", ends.Format("2006-01-02 03:04:05"))
		}
	}

	o := orm.NewOrm()
	img := make([]Image, 0)
	qs := o.QueryTable(new(Image))
	counts, _ = qs.Count()
	qs = qs.Limit(iDisplayLength, iDisplayStart)
	qs = qs.SetCond(cond)
	count, _ = qs.Count()
	qs.All(&img)
	for _, val := range img {
		var h Human
		h.Time = val.Timestamp
		h.Url = val.Url

		var d DetectedFace
		o.QueryTable(new(DetectedFace)).Filter("idimage", val.Id).One(&d)
		h.Age = d.EstimatedAge
		h.Gender = d.Gender
		human = append(human, h)
	}
	return human, count, counts
}
