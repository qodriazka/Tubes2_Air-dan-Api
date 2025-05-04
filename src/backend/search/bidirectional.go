package search

import (
	"time"
	"tubes2/utils"
)

// BidirectionalSearch performs bidirectional search from start to target in the graph
func BidirectionalSearch(g *utils.Graph, start, target string) (path []string, visited int, duration time.Duration) {
	t0 := time.Now()
	visited = 0

	forwardQueue := [][]string{{start}}
	backwardQueue := [][]string{{target}}

	visitedFromStart := map[string]bool{start: true}
	visitedFromTarget := map[string]bool{target: true}

	parentsFromStart := map[string]string{start: ""}
	parentsFromTarget := map[string]string{target: ""}

	for len(forwardQueue) > 0 && len(backwardQueue) > 0 {
		if len(forwardQueue) > 0 {
			currentPath := forwardQueue[0]
			forwardQueue = forwardQueue[1:]
			node := currentPath[len(currentPath)-1]
			visited++
			if visitedFromTarget[node] {
				return reconstructPath(parentsFromStart, parentsFromTarget, node), visited, time.Since(t0)
			}
			for _, neighbor := range g.Adj[node] {
				if !visitedFromStart[neighbor] {
					visitedFromStart[neighbor] = true
					newPath := append([]string{}, currentPath...)
					newPath = append(newPath, neighbor)
					parentsFromStart[neighbor] = node
					forwardQueue = append(forwardQueue, newPath)
				}
			}
		}
		if len(backwardQueue) > 0 {
			currentPath := backwardQueue[0]
			backwardQueue = backwardQueue[1:]
			node := currentPath[len(currentPath)-1]
			visited++

			// If the node is already visited by the forward search, path is found
			if visitedFromStart[node] {
				return reconstructPath(parentsFromStart, parentsFromTarget, node), visited, time.Since(t0)
			}
			for _, neighbor := range g.Adj[node] {
				if !visitedFromTarget[neighbor] {
					visitedFromTarget[neighbor] = true
					newPath := append([]string{}, currentPath...)
					newPath = append(newPath, neighbor)
					parentsFromTarget[neighbor] = node
					backwardQueue = append(backwardQueue, newPath)
				}
			}
		}
	}

	duration = time.Since(t0)
	return nil, visited, duration
}

func reconstructPath(parentsFromStart, parentsFromTarget map[string]string, meetingPoint string) []string {
	pathFromStart := []string{}
	for node := meetingPoint; node != ""; node = parentsFromStart[node] {
		pathFromStart = append([]string{node}, pathFromStart...)
	}
	pathFromTarget := []string{}
	for node := meetingPoint; node != ""; node = parentsFromTarget[node] {
		pathFromTarget = append(pathFromTarget, node)
	}
	pathFromTarget = pathFromTarget[1:]
	// Combine the two paths to form the final path
	return append(pathFromStart, pathFromTarget...)
}
