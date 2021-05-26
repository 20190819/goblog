package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
}

func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		return ""
	} else {
		return url.String()
	}
}

func GetRouteVariable(key string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[key]
}
