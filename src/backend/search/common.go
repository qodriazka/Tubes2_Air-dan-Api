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

