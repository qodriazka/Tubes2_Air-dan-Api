package search

import (
	"time"
	"tubes2/utils"
)

// BidirectionalSearch melakukan pencarian dua arah (forward dan backward)
func BidirectionalSearch(g *utils.Graph, target string) (TreeNode, int, time.Duration) {
	startID := g.IDs[target]
	n := len(g.AdjInt)

	// Queue dan visited untuk forward dan backward
	forwardQueue := []int{startID}
	backwardQueue := []int{}

	forwardVisited := make([]bool, n)
	backwardVisited := make([]bool, n)

	forwardVisited[startID] = true
	steps := 0
	t0 := time.Now()

	// Inisialisasi backward dari semua primitives
	for i, name := range g.Names {
		if primitives[name] {
			backwardQueue = append(backwardQueue, i)
			backwardVisited[i] = true
		}
	}

	for len(forwardQueue) > 0 && len(backwardQueue) > 0 {
		steps++

		// Expand forward
		forwardNext := []int{}
		for _, u := range forwardQueue {
			for _, v := range g.AdjInt[u] {
				if !forwardVisited[v] {
					forwardVisited[v] = true
					forwardNext = append(forwardNext, v)

					if backwardVisited[v] {
						// ditemukan titik pertemuan
						result := buildTreeNode(v, g)
						return result, steps, time.Since(t0)
					}
				}
			}
		}
		forwardQueue = forwardNext

		// Expand backward
		backwardNext := []int{}
		for _, u := range backwardQueue {
			for _, v := range g.RevInt[u] {
				if !backwardVisited[v] {
					backwardVisited[v] = true
					backwardNext = append(backwardNext, v)

					if forwardVisited[v] {
						result := buildTreeNode(v, g)
						return result, steps, time.Since(t0)
					}
				}
			}
		}
		backwardQueue = backwardNext
	}

	// Tidak ditemukan
	return TreeNode{}, steps, time.Since(t0)
}
