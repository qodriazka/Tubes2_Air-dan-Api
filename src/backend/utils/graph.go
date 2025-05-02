package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Recipe struct {
	A, B, Result string
}

type Graph struct {
	Adj map[string][]string

	AdjInt [][]int
	Names []string
	IDs map[string]int
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

func NewGraphFromJSON(jsonPath string) (*Graph, error) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON: %w", err)
	}
	var raw map[string][][2]string
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to umarshal JSON: %w", err)
	}

	g := &Graph{
		Adj:	make(map[string][]string),
		IDs:	make(map[string]int),
		Names:	[]string{},
	}

	for result, pairs := range raw {
		if _, ok := g.IDs[result]; !ok {
			g.IDs[result] = len(g.Names)
			g.Names = append(g.Names, result)
		}
		for _, p := range pairs {
			a, b := p[0], p[1]
			g.Adj[a] = append(g.Adj[a], result)
			g.Adj[b] = append(g.Adj[b], result)
			g.Adj[result] = append(g.Adj[result], a, b)
			for _, el := range[]string{a, b} {
				if _, ok := g.IDs[el]; !ok {
					g.IDs[el] = len(g.Names)
					g.Names = append(g.Names, el)
				}
			}
		}
	}

	n := len(g.Names)
	g.AdjInt = make([][]int, n)
	for res, pairs := range raw {
		rid := g.IDs[res]
		for _, p := range pairs {
			aid := g.IDs[p[0]]
			bid := g.IDs[p[1]]
			g.AdjInt[aid] = append(g.AdjInt[aid], rid)
			g.AdjInt[bid] = append(g.AdjInt[bid], rid)
			g.AdjInt[rid] = append(g.AdjInt[rid], aid, bid)
		}
	}
	return g, nil
}

func LoadAndBuildGraph(jsonPath string) (*Graph, error) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON: %w", err)
	}

	var raw map[string][][2]string
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	var recipes []Recipe
	for result, pairs := range raw {
		for _, p := range pairs {
			recipes = append(recipes, Recipe{
				A:      p[0],
				B:      p[1],
				Result: result,
			})
		}
	}

	return NewGraph(recipes), nil
}
func (g *Graph) Nodes() []string {
    keys := make([]string, 0, len(g.Adj))
    for k := range g.Adj {
        keys = append(keys, k)
    }
    return keys
}