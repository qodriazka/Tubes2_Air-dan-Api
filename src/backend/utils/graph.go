package utils

import (
    "encoding/json"
    "fmt"
    "os"
)

type Graph struct {
    // recipes raw: map[target]→[][2]string
    Recipes map[string][][2]string

    // mapping name→ID, ID→name
    IDs   map[string]int
    Names []string

    AdjInt [][]int
}

// NewGraphFromJSON membaca scraped_recipes.json dan membangun Graph
func NewGraphFromJSON(jsonPath string) (*Graph, error) {
    data, err := os.ReadFile(jsonPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read JSON: %w", err)
    }
    var raw map[string][][2]string
    if err := json.Unmarshal(data, &raw); err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    // inisialisasi graph
    g := &Graph{
        Recipes: raw,
        IDs:     make(map[string]int),
        Names:   []string{},
    }

    // assign IDs ke semua elemen (target + bahan)
    for res, pairs := range raw {
        if _, ok := g.IDs[res]; !ok {
            g.IDs[res] = len(g.Names)
            g.Names = append(g.Names, res)
        }
        for _, p := range pairs {
            a, b := p[0], p[1]
            for _, el := range []string{a, b} {
                if _, ok := g.IDs[el]; !ok {
                    g.IDs[el] = len(g.Names)
                    g.Names = append(g.Names, el)
                }
            }
        }
    }

    // bangun reverse adjacency int: dari node (res) ke bahan a,b
    n := len(g.Names)
    g.AdjInt = make([][]int, n)
    for res, pairs := range raw {
        rid := g.IDs[res]
        for _, p := range pairs {
            aid := g.IDs[p[0]]
            bid := g.IDs[p[1]]
            // res → a, res → b
            g.AdjInt[rid] = append(g.AdjInt[rid], aid, bid)
        }
    }

    return g, nil
}

// Nodes mengembalikan daftar semua elemen
func (g *Graph) Nodes() []string {
    return append([]string{}, g.Names...)
}
