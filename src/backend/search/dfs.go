package search

import (
	"time"
	"tubes2/utils"
)

func DFS(graph *utils.Graph, start, target string) (path []string, visited int, duration time.Duration) {
	t0 := time.Now()
	visited = 0
	seen := map[string]bool{}
	var dfs func(node string, currentPath []string) ([]string, bool)

	dfs = func(node string, currentPath []string) ([]string, bool) {
		visited++
		if node == target {
			return currentPath, true
		}
		seen[node] = true
		for _, neighbor := range graph.Adj[node] {
			if !seen[neighbor] {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				if result, found := dfs(neighbor, newPath); found {
					return result, true
				}
			}
		}
		return nil, false
	}

	path, _ = dfs(start, []string{start})
	duration = time.Since(t0)
	return path, visited, duration
}

func DFSInt(g *utils.Graph, startName, targetName string) ([]string, int, time.Duration) {
    startID := g.IDs[startName]
    targetID := g.IDs[targetName]

    t0 := time.Now()
    n := len(g.AdjInt)
    visited := make([]bool, n)
    parent := make([]int, n)
    for i := range parent {
        parent[i] = -1
    }

    steps := 0
    var foundPath []int
    var dfs func(u int) bool
    dfs = func(u int) bool {
        steps++
        visited[u] = true
        if u == targetID {
            // reconstruct
            var pathIDs []int
            for v := u; v != -1; v = parent[v] {
                pathIDs = append(pathIDs, v)
            }
            for i, j := 0, len(pathIDs)-1; i < j; i, j = i+1, j-1 {
                pathIDs[i], pathIDs[j] = pathIDs[j], pathIDs[i]
            }
            foundPath = pathIDs
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

    dfs(startID)
    // convert to names
    path := make([]string, len(foundPath))
    for i, id := range foundPath {
        path[i] = g.Names[id]
    }
    return path, steps, time.Since(t0)
}