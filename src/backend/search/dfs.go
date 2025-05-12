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
    pre := make(map[string][2]string)
    inPath := make(map[string]bool)

    // recursive DFS, returns true if subtree from curr can be fully built
    var dfsGraph func(curr string) bool
    dfsGraph = func(curr string) bool {
        if g.Tier(curr) == 0 {
            return true
        }
        if inPath[curr] {
            return false
        }
        inPath[curr] = true

        // gather strict recipes since enforce tier
        raw := g.RecipesFor(curr, true)
        if len(raw) == 0 {
            raw = g.RecipesFor(curr, false)
        }
        for _, combo := range raw {
            left, right := combo[0], combo[1]
            if dfsGraph(left) && dfsGraph(right) {
                pre[curr] = [2]string{left, right}
                inPath[curr] = false
                return true
            }
        }
        inPath[curr] = false
        return false
    }

    if !dfsGraph(target) {
        return nil, fmt.Errorf("no recipe path found for %q", target)
    }

    // reconstruct tree from predecessor map
    tree := BuildTreeFromPre(g, target, pre)

    // format duration and count nodes
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
