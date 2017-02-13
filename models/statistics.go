package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Statistics struct {
	Id      int       `orm:"column(id);auto" json:"id"`
	Name    string    `orm:"column(name);size(64)" json:"name"`
	Sex     string    `orm:"column(sex);size(64);null" json:"sex"`
	Ipc     string    `orm:"column(ipc);size(64)" json:"ipc"`
	Age     int       `orm:"column(age)" json:"age"`
	Created time.Time `orm:"column(created);type(datetime);index" json:"created"`
}

func (t *Statistics) TableName() string {
	return "statistics"
}

func init() {
	orm.RegisterModel(new(Statistics))
}

// GetIpcs retrieves all Ipcs matches certain condition. Returns empty list if
// no records exist
func GetIpcs() (ml []string, err error) {
	o := orm.NewOrm()

	var l []Statistics
	ipcs := make(map[string]bool)

	fields := []string{"Ipc"}
	if _, err = o.QueryTable(new(Statistics)).All(&l, fields...); err != nil {
		return nil, err
	}

	for _, v := range l {
		ipcs[v.Ipc] = true
	}

	for k, _ := range ipcs {
		ml = append(ml, k)
	}

	return
}

// AddStatistics insert a new Statistics into database and returns
// last inserted Id on success.
func AddStatistics(m *Statistics) (id int64, err error) {
	o := orm.NewOrm()
	m.Created = time.Now()
	id, err = o.Insert(m)
	return
}

// GetStatisticsById retrieves Statistics by Id. Returns error if
// Id doesn't exist
func GetStatisticsById(id int) (v *Statistics, err error) {
	o := orm.NewOrm()
	v = &Statistics{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllStatistics retrieves all Statistics matches certain condition. Returns empty list if
// no records exist
func GetAllStatistics(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, begin, end int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Statistics))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}

	// timestamp begin to end
	if begin > end {
		return nil, errors.New("Error: Invalid time. Begin must smaller than end")
	} else {
		if begin != 0 || end != 0 {
			begins := time.Unix(begin, 0)
			qs = qs.Filter("created__gte", begins.Format("2006-01-02 03:04:05"))
			ends := time.Unix(end, 0)
			qs = qs.Filter("created__lte", ends.Format("2006-01-02 03:04:05"))
		}

	}

	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Statistics
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				newtime := v.Created.Format("2006-01-02 03:04:05")
				v.Created, _ = time.Parse("2006-01-02 03:04:05", newtime)
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateStatistics updates Statistics by Id and returns error if
// the record to be updated doesn't exist
func UpdateStatisticsById(m *Statistics) (err error) {
	o := orm.NewOrm()
	v := Statistics{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteStatistics deletes Statistics by Id and returns error if
// the record to be deleted doesn't exist
func DeleteStatistics(id int) (err error) {
	o := orm.NewOrm()
	v := Statistics{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Statistics{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
