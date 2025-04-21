package router

import (
	"net/http"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/search", SearchHandler).Methods("POST")
	r.HandleFunc("/elements", ElementsHandler).Methods("GET")

	return r
}

// Implement SearchHandler dan ElementsHandler sesuai kebutuhan
