package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tsongpon/listener/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRedPlanetRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
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
		"Facebookhook",
		"GET",
		"/facebookhook",
		handler.FacebookHookGet,
	},
	Route{
		"Facebookhook",
		"POST",
		"/facebookhook",
		handler.FacebookHookPost,
	},
	Route{
		"Useractivities",
		"GET",
		"/useractivities",
		handler.QueryUserActivities,
	},
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"Login",
		"GET",
		"/login",
		handler.Login,
	},
	Route{
		"Ping",
		"GET",
		"/ping",
		handler.Ping,
	},
}
