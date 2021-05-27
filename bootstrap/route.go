package bootstrap

import (
	"github.com/gorilla/mux"
	pkgroute "github.com/yangliang4488/goblog/pkg/route"
	"github.com/yangliang4488/goblog/routes"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	pkgroute.SetRoute(router)
	routes.RegisterWebRoutes(router)
	return router
}
