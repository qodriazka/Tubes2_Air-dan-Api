package search

import (
    "time"
    "tubes2/utils"
)

// BFS mundur: mulai dari target, temukan semua bahan hingga base
// Ini mengembalikan satu jalur (chain) dari target ke elemen dasar pertama
func BFS(g *utils.Graph, target string) ([]string, int, time.Duration) {
    tid := g.IDs[target]

    t0 := time.Now()
    n := len(g.AdjInt)
    visited := make([]bool, n)
    parent := make([]int, n)
    for i := range parent {
        parent[i] = -1
    }

    queue := []int{tid}
    visited[tid] = true
    var steps int

    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        steps++

        // jika u adalah base element (tidak punya resep lagi)
        if _, has := g.Recipes[g.Names[u]]; !has {
            // reconstruct chain: dari base ke target, lalu reverse
            pathIDs := []int{}
            for v := u; v != -1; v = parent[v] {
                pathIDs = append(pathIDs, v)
            }
            // reverse to target→base
            for i, j := 0, len(pathIDs)-1; i<j; i, j = i+1, j-1 {
                pathIDs[i], pathIDs[j] = pathIDs[j], pathIDs[i]
            }
            // map to names
            path := make([]string, len(pathIDs))
            for i, id := range pathIDs {
                path[i] = g.Names[id]
            }
            return path, steps, time.Since(t0)
        }

        // expand ke bahan
        for _, v := range g.AdjInt[u] {
            if !visited[v] {
                visited[v] = true
                parent[v] = u
                queue = append(queue, v)
            }
        }
    }

    return nil, steps, time.Since(t0)
}
