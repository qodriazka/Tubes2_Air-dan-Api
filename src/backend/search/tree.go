package search

type Node struct {
    Name     string    `json:"name"`
    Combines []Combine `json:"combines"` // child pairs
}

type Combine struct {
    Left  *Node `json:"left"`
    Right *Node `json:"right"`
}

func BuildRecipeTree(target string, recipes map[string][][2]string) *Node {
    memo := make(map[string]*Node)
    return buildNode(target, recipes, memo)
}

func buildNode(name string, recipes map[string][][2]string, memo map[string]*Node) *Node {
    if n, ok := memo[name]; ok {
        return n
    }
    node := &Node{Name: name}
    memo[name] = node
    for _, pair := range recipes[name] {
        left := buildNode(pair[0], recipes, memo)
        right := buildNode(pair[1], recipes, memo)
        node.Combines = append(node.Combines, Combine{Left: left, Right: right})
    }
    return node
}
