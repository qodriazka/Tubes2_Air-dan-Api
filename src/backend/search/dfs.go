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
