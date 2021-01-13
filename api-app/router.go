package main

import (
	"github.com/gorilla/mux"
)


func RouterApp() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes{
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
			
	}

	return router
}

var routes = Routes {
	Route{"Index", "GET", "/customer/{id}", CustomerInfo},
	Route{"Index", "GET", "/customer", CustomerList},
	Route{"Index", "POST", "/customer", CustomerAdd},
	Route{"Index", "PUT", "/customer/{id}", CustomerUpdate},
	Route{"Index", "DELETE", "/customer/{id}", CustomerDelete},
}