package api

import (
	mw "go_api/middleware"
	"net/http"
	"github.com/gorilla/mux"
	hd "go_api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
			var handler http.Handler

			handler = route.HandlerFunc
			handler = mw.LoggerMiddleware(handler, route.Name)
			router.
					Methods(route.Method).
					Path(route.Pattern).
					Name(route.Name).
					Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
			"Index",
			"GET",
			"/",
			hd.GetAllPokemons,
	},
}
