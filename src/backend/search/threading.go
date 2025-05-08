package search

import (
    "fmt"
    "sync"
    "tubes2/utils"
)

// Base elements yang tidak punya resep lagi (leaf node)
var baseElements = map[string]bool{
    "Water": true,
    "Fire": true,
    "Earth": true,
    "Air": true,
}

// FindMultiplePaths mencari beberapa resep unik dari target hingga base element
func FindMultiplePaths(g *utils.Graph, target, algo string, max int) [][]string {
    switch algo {
    case "dfs":
        return findMultipleDFS(g, target, max)
    case "bfs":
        return findMultipleBFS(g, target, max)
    default:
        return findMultipleBFS(g, target, max)
    }
}

// findMultipleDFS menggunakan DFS dari target → base dengan multithread
func findMultipleDFS(g *utils.Graph, target string, max int) [][]string {
    tid := g.IDs[target]

    var mu sync.Mutex
    seen := make(map[string]bool)
    results := make([][]string, 0, max)
    var wg sync.WaitGroup

    var dfs func(path []int, visited map[int]bool)
    dfs = func(path []int, visited map[int]bool) {
        curr := path[len(path)-1]

        if baseElements[g.Names[curr]] {
            key := fmt.Sprintf("%q", idsToNames(path, g.Names))
            mu.Lock()
            if !seen[key] && len(results) < max {
                results = append(results, idsToNames(path, g.Names))
                seen[key] = true
            }
            mu.Unlock()
            return
        }

        for i := 0; i+1 < len(g.AdjInt[curr]); i += 2 {
            a, b := g.AdjInt[curr][i], g.AdjInt[curr][i+1]
            if visited[a] || visited[b] {
                continue
            }
            visited[a], visited[b] = true, true
            dfs(append(path, a, b), visited)
            visited[a], visited[b] = false, false
            if len(results) >= max {
                return
            }
        }
    }

    for i := 0; i+1 < len(g.AdjInt[tid]); i += 2 {
        a, b := g.AdjInt[tid][i], g.AdjInt[tid][i+1]
        wg.Add(1)
        go func(x, y int) {
            defer wg.Done()
            visited := map[int]bool{tid: true, x: true, y: true}
            dfs([]int{tid, x, y}, visited)
        }(a, b)
    }

    wg.Wait()
    return results
}

// findMultipleBFS menggunakan BFS dari target → base (unrolled combination)
func findMultipleBFS(g *utils.Graph, target string, max int) [][]string {
    tid := g.IDs[target]
    queue := [][]int{{tid}}

//    var mu sync.Mutex
    seen := make(map[string]bool)
    results := make([][]string, 0, max)

    for len(queue) > 0 && len(results) < max {
        path := queue[0]
        queue = queue[1:]
        curr := path[len(path)-1]

        if baseElements[g.Names[curr]] {
            key := fmt.Sprintf("%q", idsToNames(path, g.Names))
            if !seen[key] {
                seen[key] = true
                results = append(results, idsToNames(path, g.Names))
            }
            continue
        }

        for i := 0; i+1 < len(g.AdjInt[curr]); i += 2 {
            a, b := g.AdjInt[curr][i], g.AdjInt[curr][i+1]
            if containsID(path, a) || containsID(path, b) {
                continue
            }
            next := append([]int{}, path...)
            next = append(next, a, b)
            queue = append(queue, next)
        }
    }

    return results
}



func containsID(path []int, x int) bool {
    for _, v := range path {
        if v == x {
            return true
        }
    }
    return false
}

func idsToNames(path []int, names []string) []string {
    out := make([]string, len(path))
    for i, id := range path {
        out[i] = names[id]
    }
    return out
}
