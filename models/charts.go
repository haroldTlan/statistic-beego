package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type DetectedFace struct {
	Id           int     `orm:"column(iddetected_face);auto"`
	Idimage      *Image  `orm:"column(idimage);rel(fk)"`
	BoxX1        uint    `orm:"column(box_x1)"`
	BoxY1        uint    `orm:"column(box_y1)"`
	BoxX2        uint    `orm:"column(box_x2)"`
	BoxY2        uint    `orm:"column(box_y2)"`
	EstimatedAge int     `orm:"column(estimated_age);null"`
	Gender       string  `orm:"column(gender);size(10);null"`
	EyeX1        uint    `orm:"column(eye_x1);null"`
	EyeY1        uint    `orm:"column(eye_y1);null"`
	EyeX2        uint    `orm:"column(eye_x2);null"`
	EyeY2        uint    `orm:"column(eye_y2);null"`
	NoseX        uint    `orm:"column(nose_x);null"`
	NoseY        uint    `orm:"column(nose_y);null"`
	MouthX1      uint    `orm:"column(mouth_x1);null"`
	MouthY1      uint    `orm:"column(mouth_y1);null"`
	MouthX2      uint    `orm:"column(mouth_x2);null"`
	MouthY2      uint    `orm:"column(mouth_y2);null"`
	Idperson     int     `orm:"column(idperson);null"`
	Confidence   float32 `orm:"column(confidence);null"`
	HeadRotation float32 `orm:"column(head_rotation);null"`
}

type Camera struct {
	Camid string `orm:"column(camid);pk"`
}

type Image struct {
	Id        int       `orm:"column(idimage);auto"`
	Camid     *Camera   `orm:"column(camid);rel(fk)"`
	Url       string    `orm:"column(url)"`
	Timestamp time.Time `orm:"column(timestamp);type(datetime)"`
}

type ChartT struct {
	NumAge        map[string]int
	NumTime       map[int]int
	NumTimeMale   map[int]int
	NumTimeFemale map[int]int
}

type Chart struct {
	TotalMale   int             `json:"total_male"`
	TotalFemale int             `json:"total_female"`
	NumAge      [][]interface{} `json:"num_age"`
	NumMale     []int           `json:"num_male"`
	NumFemale   []int           `json:"num_female"`
	Num         []int           `json:"num"`
	NumXyMale   [][]int64       `json:"num_xy_male"`
	NumXyFemale [][]int64       `json:"num_xy_female"`
}

func init() {
	orm.RegisterModel(new(Image), new(DetectedFace), new(Camera))
}

func Charts(start, end int) Chart {
	var chart Chart
	var chartT ChartT
	chartT.NumAge = make(map[string]int)
	chartT.NumTime = make(map[int]int)
	chartT.NumTimeMale = make(map[int]int)
	chartT.NumTimeFemale = make(map[int]int)
	chart.NumAge = make([][]interface{}, 0)
	chart.Num = make([]int, 8)
	chart.NumMale = make([]int, 8)
	chart.NumFemale = make([]int, 8)
	chart.NumXyMale = make([][]int64, 0, 1)

	chart.NumXyFemale = make([][]int64, 0, 1)

	o := orm.NewOrm()
	starts := time.Unix(int64(start), 0)
	ends := time.Unix(int64(end), 0)
	img := make([]Image, 0)
	qs := o.QueryTable(new(Image)).Filter("timestamp__gte", starts.Format("2006-01-02 03:04:05")).Filter("timestamp__lte", ends.Format("2006-01-02 03:04:05"))
	if _, err := qs.Limit(100000).All(&img); err != nil {
		AddLog(err)
	}

	for _, val := range img {
		var d DetectedFace
		if err := o.QueryTable(new(DetectedFace)).Filter("idimage", val.Id).One(&d); err != nil {
			AddLog(err)
		}

		_Charts(val.Timestamp, d, &chart, &chartT)
	}
	for k, v := range chartT.NumAge {
		chart.NumAge = append(chart.NumAge, []interface{}{k, v})
	}
	for k, v := range chartT.NumTime {
		chart.Num[k] = v
	}
	for k, v := range chartT.NumTimeMale {
		chart.NumMale[k] = v
	}
	for k, v := range chartT.NumTimeFemale {
		chart.NumFemale[k] = v
	}
	return chart
}

func _Chartss(t time.Time, d DetectedFace, c *Chart, ct *ChartT) {
	t0, _ := time.Parse("15:04:05", "00:00:00")
	t3, _ := time.Parse("15:04:05", "02:59:59")
	t6, _ := time.Parse("15:04:05", "05:59:59")
	t9, _ := time.Parse("15:04:05", "08:59:59")
	t12, _ := time.Parse("15:04:05", "11:59:59")
	t15, _ := time.Parse("15:04:05", "14:59:59")
	t18, _ := time.Parse("15:04:05", "17:59:59")
	t21, _ := time.Parse("15:04:05", "20:59:59")
	t24, _ := time.Parse("15:04:05", "23:59:59")
	in, _ := time.Parse("15:04:05", t.Format("15:04:05"))

	age := d.EstimatedAge
	switch {
	case (age <= 9):
		ct.NumAge["1-9"] += 1
	case (age <= 19 && age > 9):
		ct.NumAge["10-19"] += 1
	case (age <= 29 && age > 19):
		ct.NumAge["20-29"] += 1
	case (age <= 39 && age > 29):
		ct.NumAge["30-39"] += 1
	case (age <= 49 && age > 39):
		ct.NumAge["40-49"] += 1
	case (age <= 59 && age > 49):
		ct.NumAge["50-59"] += 1
	case (age <= 69 && age > 59):
		ct.NumAge["60-69"] += 1
	case (age <= 79 && age > 69):
		ct.NumAge["70-79"] += 1
	case (age <= 89 && age > 79):
		ct.NumAge["80-89"] += 1
	case (age <= 99 && age > 89):
		ct.NumAge["90-99"] += 1
	default:
		ct.NumAge["99以上"] += 1
	}
	switch {
	case (inTimeSpan(t0, t3, in) && d.Gender == "male"):
		ct.NumTime[0] += 1
		c.NumXyMale = append(c.NumXyMale, []int64{t.Unix() * 1000, int64(age)})
		c.TotalMale += 1
		ct.NumTimeMale[0] += 1
	case (inTimeSpan(t0, t3, in) && d.Gender == "female"):
		ct.NumTime[0] += 1
		c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix() * 1000, int64(age)})
		ct.NumTimeFemale[0] += 1
		c.TotalFemale += 1

	case (inTimeSpan(t3, t6, in) && d.Gender == "male"):
		ct.NumTime[1] += 1
		c.NumXyMale = append(c.NumXyMale, []int64{t.Unix() * 1000, int64(age)})
		c.TotalMale += 1
		ct.NumTimeMale[1] += 1
	case (inTimeSpan(t3, t6, in) && d.Gender == "female"):
		ct.NumTime[1] += 1
		c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix() * 1000, int64(age)})
		ct.NumTimeFemale[1] += 1
		c.TotalFemale += 1

	case (inTimeSpan(t6, t9, in) && d.Gender == "female"):
		ct.NumTime[2] += 1
		c.NumXyMale = append(c.NumXyMale, []int64{t.Unix() * 1000, int64(age)})
		c.TotalMale += 1
		ct.NumTimeMale[2] += 1
	case (inTimeSpan(t6, t9, in) && d.Gender == "male"):
		ct.NumTime[2] += 1
		c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix() * 1000, int64(age)})
		ct.NumTimeFemale[2] += 1
		c.TotalFemale += 1

	case (inTimeSpan(t9, t12, in) && d.Gender == "male"):
		ct.NumTime[3] += 1
		c.NumXyMale = append(c.NumXyMale, []int64{t.Unix() * 1000, int64(age)})
		c.TotalMale += 1
		ct.NumTimeMale[3] += 1
	case (inTimeSpan(t9, t12, in) && d.Gender == "female"):
		ct.NumTime[3] += 1
		c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix() * 1000, int64(age)})
		ct.NumTimeFemale[3] += 1
		c.TotalFemale += 1

	case (inTimeSpan(t12, t15, in) && d.Gender == "male"):
		ct.NumTime[4] += 1
		c.NumXyMale = append(c.NumXyMale, []int64{t.Unix() * 1000, int64(age)})
		c.TotalMale += 1
		ct.NumTimeMale[4] += 1
	case (inTimeSpan(t12, t15, in) && d.Gender == "female"):
		ct.NumTime[4] += 1
		c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix() * 1000, int64(age)})
		ct.NumTimeFemale[4] += 1
		c.TotalFemale += 1

	case (inTimeSpan(t15, t18, in) && d.Gender == "male"):
		ct.NumTime[5] += 1
		c.NumXyMale = append(c.NumXyMale, []int64{t.Unix() * 1000, int64(age)})
		c.TotalMale += 1
		ct.NumTimeMale[5] += 1
	case (inTimeSpan(t15, t18, in) && d.Gender == "female"):
		ct.NumTime[5] += 1
		c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix() * 1000, int64(age)})
		ct.NumTimeFemale[5] += 1
		c.TotalFemale += 1

	case (inTimeSpan(t18, t21, in) && d.Gender == "male"):
		ct.NumTime[6] += 1
		c.NumXyMale = append(c.NumXyMale, []int64{t.Unix() * 1000, int64(age)})
		c.TotalMale += 1
		ct.NumTimeMale[6] += 1
	case (inTimeSpan(t18, t21, in) && d.Gender == "female"):
		ct.NumTime[6] += 1
		c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix() * 1000, int64(age)})
		ct.NumTimeFemale[6] += 1
		c.TotalFemale += 1

	case (inTimeSpan(t21, t24, in) && d.Gender == "male"):
		ct.NumTime[7] += 1
		c.NumXyMale = append(c.NumXyMale, []int64{t.Unix() * 1000, int64(age)})
		c.TotalMale += 1
		ct.NumTimeMale[7] += 1
	case (inTimeSpan(t21, t24, in) && d.Gender == "female"):
		ct.NumTime[7] += 1
		c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix() * 1000, int64(age)})
		ct.NumTimeFemale[7] += 1
		c.TotalFemale += 1
	default:
		fmt.Println(in)

	}
}

func _Charts(t time.Time, d DetectedFace, c *Chart, ct *ChartT) {
	t0, _ := time.Parse("15:04:05", "00:00:00")
	t3, _ := time.Parse("15:04:05", "02:59:59")
	t6, _ := time.Parse("15:04:05", "05:59:59")
	t9, _ := time.Parse("15:04:05", "08:59:59")
	t12, _ := time.Parse("15:04:05", "11:59:59")
	t15, _ := time.Parse("15:04:05", "14:59:59")
	t18, _ := time.Parse("15:04:05", "17:59:59")
	t21, _ := time.Parse("15:04:05", "20:59:59")
	t24, _ := time.Parse("15:04:05", "23:59:59")
	in, _ := time.Parse("15:04:05", t.Format("15:04:05"))

	age := d.EstimatedAge
	switch {
	case (age <= 9):
		ct.NumAge["1-9"] += 1
	case (age <= 19 && age > 9):
		ct.NumAge["10-19"] += 1
	case (age <= 29 && age > 19):
		ct.NumAge["20-29"] += 1
	case (age <= 39 && age > 29):
		ct.NumAge["30-39"] += 1
	case (age <= 49 && age > 39):
		ct.NumAge["40-49"] += 1
	case (age <= 59 && age > 49):
		ct.NumAge["50-59"] += 1
	case (age <= 69 && age > 59):
		ct.NumAge["60-69"] += 1
	case (age <= 79 && age > 69):
		ct.NumAge["70-79"] += 1
	case (age <= 89 && age > 79):
		ct.NumAge["80-89"] += 1
	case (age <= 99 && age > 89):
		ct.NumAge["90-99"] += 1
	default:
	}
	switch {
	case (inTimeSpan(t0, t3, in)):
		ct.NumTime[0] += 1
		if d.Gender == "male" {
			c.NumXyMale = append(c.NumXyMale, []int64{t.Unix(), int64(age)})
			c.TotalMale += 1
			ct.NumTimeMale[0] += 1
		} else {
			c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix(), int64(age)})
			ct.NumTimeFemale[0] += 1
			c.TotalFemale += 1
		}
	case (inTimeSpan(t3, t6, in)):
		ct.NumTime[1] += 1
		if d.Gender == "male" {
			c.NumXyMale = append(c.NumXyMale, []int64{t.Unix(), int64(age)})
			c.TotalMale += 1
			ct.NumTimeMale[1] += 1
		} else {
			c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix(), int64(age)})
			ct.NumTimeFemale[1] += 1
			c.TotalFemale += 1
		}
	case (inTimeSpan(t6, t9, in)):
		ct.NumTime[2] += 1
		if d.Gender == "male" {
			c.NumXyMale = append(c.NumXyMale, []int64{t.Unix(), int64(age)})
			c.TotalMale += 1
			ct.NumTimeMale[2] += 1
		} else {
			c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix(), int64(age)})
			ct.NumTimeFemale[2] += 1
			c.TotalFemale += 1
		}
	case (inTimeSpan(t9, t12, in)):
		ct.NumTime[3] += 1
		if d.Gender == "male" {
			c.NumXyMale = append(c.NumXyMale, []int64{t.Unix(), int64(age)})
			c.TotalMale += 1
			ct.NumTimeMale[3] += 1
		} else {
			c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix(), int64(age)})
			ct.NumTimeFemale[3] += 1
			c.TotalFemale += 1
		}
	case (inTimeSpan(t12, t15, in)):
		ct.NumTime[4] += 1
		if d.Gender == "male" {
			c.NumXyMale = append(c.NumXyMale, []int64{t.Unix(), int64(age)})
			c.TotalMale += 1
			ct.NumTimeMale[4] += 1
		} else {
			c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix(), int64(age)})
			ct.NumTimeFemale[4] += 1
			c.TotalFemale += 1
		}
	case (inTimeSpan(t15, t18, in)):
		ct.NumTime[5] += 1
		if d.Gender == "male" {
			c.NumXyMale = append(c.NumXyMale, []int64{t.Unix(), int64(age)})
			c.TotalMale += 1
			ct.NumTimeMale[5] += 1
		} else {
			c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix(), int64(age)})
			ct.NumTimeFemale[5] += 1
			c.TotalFemale += 1
		}
	case (inTimeSpan(t18, t21, in)):
		ct.NumTime[6] += 1
		if d.Gender == "male" {
			c.NumXyMale = append(c.NumXyMale, []int64{t.Unix(), int64(age)})
			c.TotalMale += 1
			ct.NumTimeMale[6] += 1
		} else {
			c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix(), int64(age)})
			ct.NumTimeFemale[6] += 1
			c.TotalFemale += 1
		}
	case (inTimeSpan(t21, t24, in)):
		ct.NumTime[7] += 1
		if d.Gender == "male" {
			c.NumXyMale = append(c.NumXyMale, []int64{t.Unix(), int64(age)})
			c.TotalMale += 1
			ct.NumTimeMale[7] += 1
		} else {
			c.NumXyFemale = append(c.NumXyFemale, []int64{t.Unix(), int64(age)})
			ct.NumTimeFemale[7] += 1
			c.TotalFemale += 1
		}
	default:
		fmt.Println(in)

	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end) || start == check || end == check
}
