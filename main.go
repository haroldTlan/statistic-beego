package main

import (
	_ "statistics/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:passwd@tcp(127.0.0.1:3306)/sam?charset=utf8&loc=Local")
	logs.SetLogger(logs.AdapterFile, `{"filename":"Statistics.log","maxdays":1}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/image", "views/image")
	beego.Run()
}
