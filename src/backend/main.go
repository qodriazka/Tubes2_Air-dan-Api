package main

import (
	"log"
	"tubes2/router"
	"tubes2/scraper"
	"tubes2/utils"
)

func main() {
	// Scrape data terbaru ke scraped_recipes.json
	if err := scraper.ScrapeToFile("scraped_recipes.json"); err != nil {
		log.Fatalf("Scraping failed: %v", err)
	}

	// Load graph dari file scraped_recipes.json
	graph, err := utils.NewGraph("scraped_recipes.json")
	if err != nil {
		log.Fatalf("Error loading graph: %v", err)
	}

	// Siapkan router dengan dependency graph
	r := router.SetupRouter(graph)

	// Jalankan HTTP server di port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
