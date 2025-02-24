package router

import (
	"playsee/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/test-1", handlers.Test1).Methods("POST")

	return router
}
