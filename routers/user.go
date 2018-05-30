package routers

import (
	"golaman-pack/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) *mux.Router {
	router.Handle("/v1/user/create",
		negroni.New(
			negroni.HandlerFunc(controllers.CreateUser),
		)).Methods("POST")
	return router
}