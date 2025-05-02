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

func BFSInt(g *utils.Graph, startName, targetName string) ([]string, int, time.Duration) {
    // konversi nama ke ID
    startID := g.IDs[startName]
    targetID := g.IDs[targetName]

    t0 := time.Now()
    n := len(g.AdjInt)
    visited := make([]bool, n)
    parent := make([]int, n)
    for i := range parent {
        parent[i] = -1
    }

    queue := []int{startID}
    visited[startID] = true
    steps := 0

    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        steps++

        if u == targetID {
            // reconstruct path of IDs
            pathIDs := []int{}
            for v := u; v != -1; v = parent[v] {
                pathIDs = append(pathIDs, v)
            }
            // reverse
            for i, j := 0, len(pathIDs)-1; i < j; i, j = i+1, j-1 {
                pathIDs[i], pathIDs[j] = pathIDs[j], pathIDs[i]
            }
            // convert to names
            path := make([]string, len(pathIDs))
            for i, id := range pathIDs {
                path[i] = g.Names[id]
            }
            return path, steps, time.Since(t0)
        }

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