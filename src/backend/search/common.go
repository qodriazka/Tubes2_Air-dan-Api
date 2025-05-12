package search

import "tubes2/utils"

// Node represents one node in the recipe tree
type Node struct {
	Name     string  `json:"name"`
	Combines []*Node `json:"combines,omitempty"`
}

// SearchResult holds a found recipe tree and stats
type SearchResult struct {
	Recipe       *Node  `json:"recipe"`
	NodesVisited int    `json:"nodesVisited"`
	Duration     string `json:"duration"` // e.g. "15ms"
}

// cloneNode deep-clones a tree node for safe branching
func cloneNode(n *Node) *Node {
	if n == nil {
		return nil
	}
	copy := &Node{Name: n.Name}
	for _, c := range n.Combines {
		copy.Combines = append(copy.Combines, cloneNode(c))
	}
	return copy
}

// allLeavesAreBase checks if all leaves in the tree are base elements
// uses provided predicate isBase to test leaf nodes
func allLeavesAreBase(n *Node, isBase func(string) bool) bool {
	if len(n.Combines) == 0 {
		return isBase(n.Name)
	}
	for _, c := range n.Combines {
		if !allLeavesAreBase(c, isBase) {
			return false
		}
	}
	return true
}

func BuildTreeFromPre(g *utils.Graph, target string, pre map[string][2]string) *Node {
	root := &Node{Name: target}
	var build func(n *Node)
	build = func(n *Node) {
		pair, ok := pre[n.Name]
		if !ok {
			// tidak ada resep → n adalah leaf/base
			return
		}
		left := &Node{Name: pair[0]}
		right := &Node{Name: pair[1]}
		n.Combines = []*Node{left, right}
		build(left)
		build(right)
	}
	build(root)
	return root
}

// findLeaf returns the first non-base leaf for expansion
func findLeaf(n *Node, g *utils.Graph) *Node {
	if n == nil {
		return nil
	}
	var result *Node
	var dfs func(*Node)
	dfs = func(curr *Node) {
		if result != nil {
			return
		}
		if len(curr.Combines) == 0 && g.Tier(curr.Name) > 0 {
			result = curr
			return
		}
		for _, c := range curr.Combines {
			dfs(c)
		}
	}
	dfs(n)
	return result
}

// findLeafInTree finds leaf by name in a cloned tree
func findLeafInTree(n *Node, name string) *Node {
	if n == nil {
		return nil
	}
	if n.Name == name && len(n.Combines) == 0 {
		return n
	}
	for _, c := range n.Combines {
		if found := findLeafInTree(c, name); found != nil {
			return found
		}
	}
	return nil
}

// validCombo ensures all ingredients have tier < currentTier
func validCombo(combo []string, currentTier int, g *utils.Graph) bool {
	for _, ing := range combo {
		if g.Tier(ing) >= currentTier {
			return false
		}
	}
	return true
}
