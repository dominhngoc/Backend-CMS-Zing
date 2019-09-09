package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func SetStatic() {
	beego.SetStaticPath("/storage", "storage/")
}

func Cors() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
