package main

import (
    "encoding/json"
    "log"
    "os"

    "tubes2/router"
    "tubes2/utils"
)

func main() {
    // Load raw recipes map
    raw := make(map[string][][2]string)
    data, err := os.ReadFile("scraped_recipes.json")
    if err != nil {
        log.Fatalf("failed to read scraped_recipes.json: %v", err)
    }
    if err := json.Unmarshal(data, &raw); err != nil {
        log.Fatalf("failed to unmarshal JSON: %v", err)
    }

    // Build graph from same JSON
    g, err := utils.NewGraphFromJSON("scraped_recipes.json")
    if err != nil {
        log.Fatalf("failed to build graph: %v", err)
    }

    // Start router with both graph and raw recipes
    r := router.SetupRouter(g, raw)
    log.Println("listening on :8080")
    log.Fatal(r.Run(":8080"))
}
