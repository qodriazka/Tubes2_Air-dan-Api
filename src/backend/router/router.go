package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/search", SearchHandler).Methods("GET")
	r.HandleFunc("/elements", ElementsHandler).Methods("GET")
	r.Use(middlewareCORS)
	return r
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	element := r.URL.Query().Get("element")
	res := map[string][]string{
		"recipes": {element + "_dummy1", element + "_dummy2"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func ElementsHandler(w http.ResponseWriter, r *http.Request) {
	elements := []string{"air", "fire", "earth", "water"}
	json.NewEncoder(w).Encode(map[string][]string{"elements": elements})
}

func middlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
