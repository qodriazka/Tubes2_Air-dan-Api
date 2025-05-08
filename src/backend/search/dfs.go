package search

import (
    "time"
    "tubes2/utils"
)

// DFS mundur: rekursif trace sampai base element
func DFS(g *utils.Graph, target string) ([]string, int, time.Duration) {
    tid := g.IDs[target]

    t0 := time.Now()
    n := len(g.AdjInt)
    visited := make([]bool, n)
    parent := make([]int, n)
    for i := range parent {
        parent[i] = -1
    }

    var steps int
    var foundID int = -1

    var dfs func(u int) bool
    dfs = func(u int) bool {
        steps++
        visited[u] = true
        name := g.Names[u]
        // leaf?
        if _, has := g.Recipes[name]; !has {
            foundID = u
            return true
        }
        for _, v := range g.AdjInt[u] {
            if !visited[v] {
                parent[v] = u
                if dfs(v) {
                    return true
                }
            }
        }
        return false
    }

    dfs(tid)

    if foundID < 0 {
        return nil, steps, time.Since(t0)
    }
    // reconstruct and reverse
    pathIDs := []int{}
    for v := foundID; v != -1; v = parent[v] {
        pathIDs = append(pathIDs, v)
    }
    for i, j := 0, len(pathIDs)-1; i<j; i, j = i+1, j-1 {
        pathIDs[i], pathIDs[j] = pathIDs[j], pathIDs[i]
    }
    path := make([]string, len(pathIDs))
    for i, id := range pathIDs {
        path[i] = g.Names[id]
    }
    return path, steps, time.Since(t0)
}
