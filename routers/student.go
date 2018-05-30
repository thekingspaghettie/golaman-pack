package routers

import (
	"golaman-pack/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func StudentRoutes(router *mux.Router) *mux.Router {
	router.Handle("/v1/student/read",
		negroni.New(
			negroni.HandlerFunc(controllers.ReadStudents),
		)).Methods("GET")
	router.Handle("/v1/student/create",
		negroni.New(
			negroni.HandlerFunc(controllers.CreateStudent),
		)).Methods("POST")
	return router
}