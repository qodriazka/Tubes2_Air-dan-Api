package main

import (
	"log"
	"net/http"
	"tubes2/router"
)

func main() {
	r := router.SetupRouter()
	port := ":8080"
	log.Printf("Starting server on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
