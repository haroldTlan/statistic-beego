package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

// GetAllStatistics retrieves all Statistics matches certain condition. Returns empty list if
// no records exist
func AllStatistics(query map[string]string, sortby []string, order []string,
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
	if _, err = qs.Limit(limit, offset).All(&l); err == nil {
		for _, v := range l {
			newtime := v.Created.Format("2006-01-02 03:04:05")
			v.Created, _ = time.Parse("2006-01-02 03:04:05", newtime)
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}
