package main

import (
    "fmt"
    "log"
    "tubes2/router"
    "tubes2/utils"
)

func main() {
    // Load graph dari JSON
    graph, err := utils.NewGraphFromJSON("scraped_recipes.json")
    if err != nil {
        log.Fatalf("Failed to load graph: %v", err)
    }

    // Setup dan jalankan router
    r := router.SetupRouter(graph)

    // Jalankan di port 8080 (atau sesuaikan)
    port := ":8080"
    fmt.Printf("Server running on http://localhost%s\n", port)
    if err := r.Run(port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
