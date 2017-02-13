package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"statistics/models"
	"strconv"
	"strings"
	"time"
)

// SearchController operations for searching  Statistics
type SearchController struct {
	beego.Controller
}

type Datatable struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// URLMapping ...
func (c *SearchController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)

}

//需要搞懂一些问题
//以下的方式可以更加简洁，可以直接使用datatable的search，可是前端的问题，导致需要引入新的3个参数来控制筛选
//还有一个以待日后解决的问题就是，过去由于主动权的问题，导致开发都很混乱，get post，不得已按照前端的要求来使用，即使是错的，
//以后如果没有特定的分歧，一定要严格按照每种方法的意义来使用。
//以下，就是由于前端的分歧而导致的错误用法，理应是get方法

// Post ...
// @Title Post
// @Description Pagination for Datatables
// @Param   body        body    models.AoData   true        "body for aoData: ipc=1&begin=1486051200&end=1486656000&sEcho=6&iColumns=4&sColumns=%2C%2C%2C&iDisplayStart=0&iDisplayLength=10&mDataProp_0=&sSearch_0=&bRegex_0=false&bSearchable_0=false&mDataProp_1=1&sSearch_1=&bRegex_1=false&bSearchable_1=true&mDataProp_2=2&sSearch_2=&bRegex_2=false&bSearchable_2=true&mDataProp_3=3&sSearch_3=&bRegex_3=false&bSearchable_3=true&sSearch=&bRegex=false"
// @Success 201 {int} models.Statistics
// @Failure 403 body is empty
// @router / [post]
func (c *SearchController) Post() {
	fmt.Printf("%+v", string(c.Ctx.Input.RequestBody))
	aColumns := []string{
		"Sex",
		"Age",
		"Created",
	}

	maps, count, counts := models.Datatables(aColumns, new(models.Statistics), c.Ctx.Input)

	data := make(map[string]interface{}, count)
	var output = make([][]interface{}, len(maps))
	for i, m := range maps {
		for _, v := range aColumns {
			if v == "Created" {
				output[i] = append(output[i], m[v].(time.Time).Format("2006-01-02 15:04:05"))
			} else if v == "Sex" {
				if m[v] == "male" {
					output[i] = append(output[i], m[v])
					output[i] = append(output[i], "男")
				} else {
					output[i] = append(output[i], m[v])
					output[i] = append(output[i], "女")
				}
			} else {
				output[i] = append(output[i], m[v])
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

// Get ...
// @Title Get
// @Description search Statistics
// @Param   query   query   string  false   "Filter. e.g. col1:v1,col2:v2 ..."
// @Param   sortby  query   string  false   "Sorted-by fields. e.g. col1,col2 ..."
// @Param   order   query   string  false   "Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param   limit   query   string  false   "Limit the size of result set. Must be an integer"
// @Param   offset  query   string  false   "Start position of result set. Must be an integer"
// @Param   begin   query   string  false   "Start position of time. Must be an integer.TimeStamp!!"
// @Param   end     query   string  false   "End position of time. Must be an integer.TimeStamp!!"
// @Success 200 {object} models.Statistics
// @Failure 403
// @router / [get]
func (c *SearchController) Get() {
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 5
	var offset int64
	var begin int64
	var end int64

	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// begin: timestamp
	if v, err := c.GetInt64("begin"); err == nil {
		begin = v
	}
	// end: timestamp
	if v, err := c.GetInt64("end"); err == nil {
		end = v
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.AllStatistics(query, sortby, order, offset, limit, begin, end)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

func (c *SearchController) TestAjax() {
	fmt.Println("?????\n")
	fmt.Printf("%+v", c.Input())
}
