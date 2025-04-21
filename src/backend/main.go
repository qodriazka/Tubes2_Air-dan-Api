package main

import (
	"log"
	"net/http"
	"tubes2/router"
)

func main() {
	r := router.SetupRouter()
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
