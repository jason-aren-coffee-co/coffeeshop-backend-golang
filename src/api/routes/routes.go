package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Url     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func SetupRoutes(r *mux.Router) {
	// Setup order routes
	for _, route := range orderRoutes {
		r.HandleFunc(route.Url, route.Handler).Methods(route.Method)
	}
	//Set up authentication routes
	for _, route := range authRoutes {
		r.HandleFunc(route.Url, route.Handler).Methods(route.Method)
	}
}
