package main

import (
	"log"
	"net/http"
	"tubes2/router"
	"tubes2/utils"
)

func main() {
	g, err := utils.LoadAndBuildGraph("scraped_recipes.json")
	if err != nil {
	  log.Fatalf("Failed to load graph: %v", err)
	}
	r := router.SetupRouter(g)
	port := ":8080"
	log.Printf("Starting server on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
