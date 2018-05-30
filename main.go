package main

import (
	"github.com/codegangsta/negroni"
	"github.com/astaxie/beego"
	"net/http"
	"github.com/astaxie/beego/orm"

	_"github.com/go-sql-driver/mysql"

	"golaman-pack/routers"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:admin@/golaman?charset=utf8", 30)

	router 			:= routers.InitRoutes()
	projectPort 	:= beego.AppConfig.String("httpport")
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":" + projectPort, n)
}
