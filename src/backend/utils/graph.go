package utils

import (
    "encoding/json"
    "fmt"
    "os"
)

// Recipe merepresentasikan satu cara membuat target dari 2 bahan
type Recipe struct {
    Ingredients []string
}

// Element hanya menyimpan tier (kita tidak pakai Recipes di sini lagi)
type Element struct {
    Tier int
}

// Graph memetakan name → Element (untuk tier) dan
// name → daftar Recipe (untuk struktur graf)
type Graph struct {
    Elements map[string]*Element     // untuk cek Tier()
    Recipes  map[string][][]string   // raw recipe list
}

// NewGraph membaca JSON dan mengisi kedua map di atas
func NewGraph(path string) (*Graph, error) {
    raw, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed reading file: %w", err)
    }

    // Struktur per‐elemen di JSON
    var data map[string]struct {
        Tier    int        `json:"tier"`
        Recipes [][]string `json:"recipes"`
    }
    if err := json.Unmarshal(raw, &data); err != nil {
        return nil, fmt.Errorf("failed parsing JSON: %w", err)
    }

    g := &Graph{
        Elements: make(map[string]*Element),
        Recipes:  make(map[string][][]string),
    }

    // Isi kedua map
    for name, info := range data {
        g.Elements[name] = &Element{Tier: info.Tier}
        g.Recipes[name] = info.Recipes
    }

    return g, nil
}

// Tier mengembalikan tier elemen, atau -1 kalau tidak ada
func (g *Graph) Tier(name string) int {
    if e, ok := g.Elements[name]; ok {
        return e.Tier
    }
    return -1
}

// RecipesFor mengembalikan daftar resep raw
// jika enforceTier==true maka filter tier(ing)<tier(target)
func (g *Graph) RecipesFor(target string, enforceTier bool) [][]string {
    raw, ok := g.Recipes[target]
    if !ok {
        return nil
    }
    if !enforceTier {
        return raw
    }
    var out [][]string
    t0 := g.Tier(target)
    for _, combo := range raw {
        valid := true
        for _, ing := range combo {
            if g.Tier(ing) >= t0 {
                valid = false
                break
            }
        }
        if valid {
            out = append(out, combo)
        }
    }
    return out
}
