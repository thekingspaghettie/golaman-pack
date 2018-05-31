package routers

import (
	"golaman-pack/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) *mux.Router {
	router.Handle("/v1/user/guest/create",
		negroni.New(
			negroni.HandlerFunc(controllers.CreateGuest),
		)).Methods("POST")
	router.Handle("/v1/user/guest/read",
		negroni.New(
			negroni.HandlerFunc(controllers.ReadGuest),
		)).Methods("POST")
	router.Handle("/v1/user/create",
		negroni.New(
			negroni.HandlerFunc(controllers.CreateUser),
		)).Methods("POST")
	return router
}