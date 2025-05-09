package search

type Node struct {
    Name     string    `json:"name"`
    Combines []Combine `json:"combines"` // child pairs
}

type Combine struct {
    Left  *Node `json:"left"`
    Right *Node `json:"right"`
}

var primitives = map[string]bool{
    "Water": true,
    "Earth": true,
    "Fire":  true,
    "Air":   true,
}

// TreeNode adalah format yang akan dikirim ke frontend
type TreeNode struct {
    Name     string     `json:"name"`
    Combines []TreeNode `json:"combines,omitempty"`
}

func BuildRecipeTree(target string, recipes map[string][][2]string) *Node {
    visited := make(map[string]bool)
    root, _ := buildNode(target, recipes, visited)
    return root
}

func buildNode(name string, recipes map[string][][2]string, visited map[string]bool) (*Node, bool) {
    if primitives[name] {
        return &Node{Name: name}, true
    }
    if visited[name] {
        return nil, false
    }
    visited[name] = true
    defer func() { visited[name] = false }()
    node := &Node{Name: name}
    for _, pair := range recipes[name] {
        left, okL := buildNode(pair[0], recipes, visited)
        right, okR := buildNode(pair[1], recipes, visited)
        if okL && okR {
            node.Combines = append(node.Combines, Combine{Left: left, Right: right})
        }
    }
    if len(node.Combines) == 0 {
        return nil, false
    }
    return node, true
}