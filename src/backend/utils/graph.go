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