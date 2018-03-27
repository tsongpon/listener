package main

import (
	"net/http"

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
		FacebookHookGet,
	},
	Route{
		"Facebookhook",
		"POST",
		"/facebookhook",
		FacebookHookPost,
	},
	Route{
		"Useractivities",
		"GET",
		"/useractivities",
		QueryUserActivities,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Login",
		"GET",
		"/login",
		Login,
	},
	Route{
		"Ping",
		"GET",
		"/ping",
		Ping,
	},
}
