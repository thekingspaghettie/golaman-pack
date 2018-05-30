package routers
import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = StudentRoutes(router)
	router = UserRoutes(router)
	return router
}