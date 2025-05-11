package search

import (
    "fmt"
    "time"
    "tubes2/utils"
)

// SearchDFS performs a pure depth-first search on the element graph
// until it finds one complete recipe tree (all leaves are base elements).
// Returns a single SearchResult with stats and duration.
func SearchDFS(g *utils.Graph, target string) ([]SearchResult, error) {
    start := time.Now()
    // predecessor map: for each non-base element, which pair built it
    pre := make(map[string][2]string)
    // inPath tracks current recursion stack to avoid cycles
    inPath := make(map[string]bool)
    var steps int

    // recursive DFS, returns true if subtree from curr can be fully built
    var dfsGraph func(curr string) bool
    dfsGraph = func(curr string) bool {
        steps++
        if g.Tier(curr) == 0 {
            return true
        }
        if inPath[curr] {
            return false
        }
        inPath[curr] = true

        // gather recipes: strict first, then fallback
        recipes := g.RecipesFor(curr, true)
        if len(recipes) == 0 {
            recipes = g.RecipesFor(curr, false)
        }

        for _, combo := range recipes {
            left, right := combo[0], combo[1]
            if dfsGraph(left) && dfsGraph(right) {
                // only record successful path
                pre[curr] = [2]string{left, right}
                inPath[curr] = false
                return true
            }
        }

        inPath[curr] = false
        return false
    }

    found := dfsGraph(target)
    if !found {
        return nil, fmt.Errorf("no recipe path found for %q", target)
    }

    // reconstruct tree from predecessor map
    tree := BuildTreeFromPre(g, target, pre)

    // format duration to milliseconds
    dur := time.Since(start)
    duration := fmt.Sprintf("%.3fms", float64(dur.Nanoseconds())/1e6)

    return []SearchResult{{
        Recipe:       tree,
        NodesVisited: steps,
        Duration:     duration,
    }}, nil
}
