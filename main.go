package main

import (
	_ "faceStatis/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:passwd@tcp(127.0.0.1:3306)/FaceTest?charset=utf8&loc=Local")
	logs.SetLogger(logs.AdapterFile, `{"filename":"Statistics.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/Cam_Images", "/bary/face/Cam_Images")
	beego.Run()
}
