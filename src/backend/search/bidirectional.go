package search

import (
	"time"
	"tubes2/utils"
)

func BidirectionalSearch(g *utils.Graph, target string) (path []string, visited int, duration time.Duration) {
	t0 := time.Now()
	visited = 0

	tid := g.IDs[target]

	// gunakan ID base elements
	baseIDs := []int{
		g.IDs["Water"],
		g.IDs["Fire"],
		g.IDs["Earth"],
		g.IDs["Air"],
	}

	forwardQueue := [][]int{{tid}}
	visitedFromTarget := map[int]bool{tid: true}
	parentFromTarget := map[int]int{tid: -1}

	backwardQueue := [][]int{}
	visitedFromBase := map[int]bool{}
	parentFromBase := map[int]int{}

	for _, bid := range baseIDs {
		backwardQueue = append(backwardQueue, []int{bid})
		visitedFromBase[bid] = true
		parentFromBase[bid] = -1
	}

	for len(forwardQueue) > 0 && len(backwardQueue) > 0 {
		// Forward dari target
		if len(forwardQueue) > 0 {
			path := forwardQueue[0]
			forwardQueue = forwardQueue[1:]
			curr := path[len(path)-1]
			visited++

			if visitedFromBase[curr] {
				return reconstructIntPath(parentFromTarget, parentFromBase, curr, g), visited, time.Since(t0)
			}

			for i := 0; i+1 < len(g.AdjInt[curr]); i += 2 {
				a := g.AdjInt[curr][i]
				b := g.AdjInt[curr][i+1]

				for _, nbr := range []int{a, b} {
					if !visitedFromTarget[nbr] {
						visitedFromTarget[nbr] = true
						parentFromTarget[nbr] = curr
						newPath := append([]int{}, path...)     // duplikat path
						newPath = append(newPath, nbr)         // tambahkan neighbor
						forwardQueue = append(forwardQueue, newPath)
					}
				}
			}
		}

		// Backward dari base
		if len(backwardQueue) > 0 {
			path := backwardQueue[0]
			backwardQueue = backwardQueue[1:]
			curr := path[len(path)-1]
			visited++

			if visitedFromTarget[curr] {
				return reconstructIntPath(parentFromTarget, parentFromBase, curr, g), visited, time.Since(t0)
			}

			for _, nbr := range g.AdjInt[curr] {
				if !visitedFromBase[nbr] {
					visitedFromBase[nbr] = true
					parentFromBase[nbr] = curr
					newPath := append([]int{}, path...)     // duplikat path
					newPath = append(newPath, nbr)         // tambahkan neighbor
					backwardQueue = append(backwardQueue, newPath)
				}
			}
		}
	}

	return nil, visited, time.Since(t0)
}

func reconstructIntPath(parentFromTarget, parentFromBase map[int]int, meet int, g *utils.Graph) []string {
	path1 := []int{}
	for v := meet; v != -1; v = parentFromTarget[v] {
		path1 = append([]int{v}, path1...)
	}
	path2 := []int{}
	for v := parentFromBase[meet]; v != -1; v = parentFromBase[v] {
		path2 = append(path2, v)
	}
	full := append(path1, path2...)
	names := make([]string, len(full))
	for i, id := range full {
		names[i] = g.Names[id]
	}
	return names
}
