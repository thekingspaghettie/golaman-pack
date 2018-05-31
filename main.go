package main

import (
	"github.com/codegangsta/negroni"
	"github.com/astaxie/beego"
	"net/http"
	
	_"github.com/go-sql-driver/mysql"

	"golaman-pack/routers"
)

func main() {
	router 			:= routers.InitRoutes()
	projectPort 	:= beego.AppConfig.String("httpport")
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":" + projectPort, n)
}
