package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"./logger"

)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandleFunc
		handler = logger.Logger(handler,route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router

}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},Route{
		"Max",
		"GET",
		"/max",
		TheMax,
	},Route{
		"findbyId",
		"POST",
		"/find_one",
		FindById,
	},

}
