package utils

type Recipe struct {
	A, B, Result string
}

type Graph struct {
	Adj map[string][]string
}

func NewGraph(recipes []Recipe) *Graph {
	g := &Graph{Adj: make(map[string][]string)}
	for _, r := range recipes {
		g.Adj[r.A] = append(g.Adj[r.A], r.Result)
		g.Adj[r.B] = append(g.Adj[r.B], r.Result)
		g.Adj[r.Result] = append(g.Adj[r.Result], r.A, r.B)
	}
	return g
}

func (g *Graph) Nodes() []string {
	keys := make([]string, 0, len(g.Adj))
	for k := range g.Adj {
		keys = append(keys, k)
	}
	return keys
}
