// search/bfs.go
package search

import (
	"time"
	"tubes2/utils"
)

// BFS performs breadth-first search from start to target in the graph
func BFS(graph *utils.Graph, start, target string) (path []string, visited int, duration time.Duration) {
	t0 := time.Now()
	visited = 0
	queue := [][]string{{start}}
	seen := map[string]bool{start: true}

	for len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		node := currentPath[len(currentPath)-1]
		visited++

		if node == target {
			duration = time.Since(t0)
			return currentPath, visited, duration
		}

		for _, neighbor := range graph.Adj[node] {
			if !seen[neighbor] {
				seen[neighbor] = true
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}

	duration = time.Since(t0)
	return nil, visited, duration
}
