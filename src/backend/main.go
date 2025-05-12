package main

import (
    "encoding/json"
    "fmt"
    "log"

    "tubes2/search"
    "tubes2/utils"
)

func main() {
    graph, err := utils.NewGraph("scraped_recipes.json")
    if err != nil {
        log.Fatalf("Error loading graph: %v", err)
    }
    target := "Grilled cheese"         // ganti sesuai yang mau kamu tes
    maxRecipes := 3           // berapa banyak recipe yang ingin dicari

    fmt.Println("=== BIDIRECTIONAL MULTIPLE RESULTS ===")
    multi, err := search.SearchDFSMultiple(graph, target, maxRecipes)
    if err != nil {
        log.Fatalf("Bidirectional error: %v", err)
    }
    out, _ := json.MarshalIndent(multi, "", "  ")
    fmt.Println(string(out))
}
