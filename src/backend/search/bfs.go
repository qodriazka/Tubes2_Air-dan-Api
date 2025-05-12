package search

import (
	"container/list"
	"fmt"
	"time"
	"tubes2/utils"
)

// SearchBFS performs a pure breadth-first search on the element graph
// until it finds one complete recipe tree (all leaves are base elements).
// Returns a single SearchResult with stats (nodes visited and duration).
func SearchBFS(g *utils.Graph, target string) ([]SearchResult, error) {
	start := time.Now()

	// predecessor map: for each non-base element, record the recipe that builds it
	pre := make(map[string][2]string)
	visited := make(map[string]bool)

	// BFS queue over element names
	type state struct{ elem string }
	q := list.New()
	q.PushBack(state{target})
	visited[target] = true

	// BFS loop
	for q.Len() > 0 {
		curr := q.Remove(q.Front()).(state).elem

		// Skip expansion if base element
		if g.Tier(curr) == 0 {
			continue
		}

		// Try recipes with tier enforcement, fallback if none
		combos := g.RecipesFor(curr, true)
		if len(combos) == 0 {
			combos = g.RecipesFor(curr, false)
		}
		if len(combos) == 0 {
			continue
		}

		// Record first recipe for curr
		pair := combos[0]
		pre[curr] = [2]string{pair[0], pair[1]}

		// Enqueue ingredients
		for _, ing := range pair {
			if !visited[ing] {
				visited[ing] = true
				q.PushBack(state{ing})
			}
		}
	}

	// Reconstruct tree from pre map
	tree := BuildTreeFromPre(g, target, pre)

	// Format duration to milliseconds
	dur := time.Since(start)
	duration := fmt.Sprintf("%.3fms", float64(dur.Nanoseconds())/1e6)
	var countNodes func(n *Node) int
	countNodes = func(n *Node) int {
		if n == nil {
			return 0
		}
		cnt := 1
		for _, c := range n.Combines {
			cnt += countNodes(c)
		}
		return cnt
	}
	nodesVisited := countNodes(tree)
	return []SearchResult{{
		Recipe:       tree,
		NodesVisited: nodesVisited,
		Duration:     duration,
	}}, nil
}
