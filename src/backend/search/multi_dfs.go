package search

import (
	"encoding/json"
	"fmt"
	"time"
	"tubes2/utils"
)

// SearchDFSMultiple finds up to maxRecipes distinct strict recipes using DFS backtracking.
func SearchDFSMultiple(
	g *utils.Graph,
	target string,
	maxRecipes int,
) ([]SearchResult, error) {
	// only strict initial combos
	rawInit := g.RecipesFor(target, true)
	if len(rawInit) == 0 {
		return nil, fmt.Errorf("no strict initial combos for %q", target)
	}
	var initCombos [][2]string
	for _, r := range rawInit {
		if len(r) == 2 {
			initCombos = append(initCombos, [2]string{r[0], r[1]})
		}
	}

	results := []SearchResult{}
	visitedTree := make(map[string]bool)
	start := time.Now()

	// recursive DFS backtracking to collect recipes
	var dfsAll func(pre map[string][2]string, combos [][2]string)
	dfsAll = func(pre map[string][2]string, combos [][2]string) {
		if len(results) >= maxRecipes {
			return
		}
		// if complete, collect
		if allLeavesAreBase(BuildTreeFromPre(g, target, pre), func(name string) bool { return g.Tier(name) == 0 }) {
			tree := BuildTreeFromPre(g, target, pre)
			sig, _ := json.Marshal(tree)
			key := string(sig)
			if !visitedTree[key] {
				visitedTree[key] = true
				// count nodes
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
				results = append(results, SearchResult{
					Recipe:       tree,
					NodesVisited: countNodes(tree),
					Duration:     fmt.Sprintf("%.3fms", float64(time.Since(start).Nanoseconds())/1e6),
				})
				if len(results) >= maxRecipes {
					return
				}
			}
		}
		// find first expandable leaf
		var leafName string
		var findLeaf func(n *Node) bool
		findLeaf = func(n *Node) bool {
			if len(n.Combines) == 0 && g.Tier(n.Name) > 0 {
				leafName = n.Name
				return true
			}
			for _, c := range n.Combines {
				if findLeaf(c) {
					return true
				}
			}
			return false
		}
		root := BuildTreeFromPre(g, target, pre)
		if !findLeaf(root) {
			return
		}
		// try strict recipes for this leaf
		raw := g.RecipesFor(leafName, true)
		for _, combo := range raw {
			if len(combo) != 2 {
				continue
			}
			// extend pre copy
			newPre := make(map[string][2]string, len(pre))
			for k, v := range pre {
				newPre[k] = v
			}
			newPre[leafName] = [2]string{combo[0], combo[1]}
			dfsAll(newPre, initCombos)
			if len(results) >= maxRecipes {
				return
			}
		}
	}

	// seed with each initial combo
	for _, ic := range initCombos {
		if len(results) >= maxRecipes {
			break
		}
		pre := map[string][2]string{target: ic}
		dfsAll(pre, initCombos)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no strict recipes found for %q", target)
	}
	return results, nil
}
