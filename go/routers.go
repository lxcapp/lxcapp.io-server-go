package lxcapp

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
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
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/zmatsh/lxcapp/1.0.0/",
		Index,
	},

	Route{
		"LxcappContrinerContrineridConfigGet",
		"GET",
		"/zmatsh/lxcapp/1.0.0/lxcapp/Contriner/{Contrinerid}/config",
		LxcappContrinerContrineridConfigGet,
	},

	Route{
		"LxcappContrinerContrineridConfigPost",
		"POST",
		"/zmatsh/lxcapp/1.0.0/lxcapp/Contriner/{Contrinerid}/config",
		LxcappContrinerContrineridConfigPost,
	},

	Route{
		"LxcappContrinerContrineridDelete",
		"DELETE",
		"/zmatsh/lxcapp/1.0.0/lxcapp/Contriner/{Contrinerid}",
		LxcappContrinerContrineridDelete,
	},

	Route{
		"LxcappContrinerListGet",
		"GET",
		"/zmatsh/lxcapp/1.0.0/lxcapp/Contriner/list",
		LxcappContrinerListGet,
	},

	Route{
		"LxcappContrinerListPost",
		"POST",
		"/zmatsh/lxcapp/1.0.0/lxcapp/Contriner/list",
		LxcappContrinerListPost,
	},

}