package search

import (
    "fmt"
    "sync"
    "tubes2/utils"
)

// Pencarian multithreaded untuk multiple recipe
func FindMultiplePaths(g *utils.Graph, start, target, algo string, max int) [][]string {
    var mu sync.Mutex
    var results [][]string
    seen := make(map[string]bool)
    wg := sync.WaitGroup{}

    for i := 0; i < max; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            var path []string
            if algo == "dfs" {
                path, _, _ = DFS(g, start, target)
            } else {
                path, _, _ = BFS(g, start, target)
            }

            if len(path) > 0 {
                key := fmt.Sprintf("%v", path)
                mu.Lock()
                if !seen[key] {
                    results = append(results, path)
                    seen[key] = true
                }
                mu.Unlock()
            }
        }()
    }

    wg.Wait()
    return results
}
