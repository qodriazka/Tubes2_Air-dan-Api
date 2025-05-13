package search

import (
	"container/list"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"tubes2/utils"
)

// Study: BFS local, seeded dengan pre[target], tanpa mengubah g.Recipes
func SearchBFSWithPre(
	g *utils.Graph,
	target string,
	pre map[string][2]string,
) ([]SearchResult, error) {
	start := time.Now()

	// Ambil dua bahan awal dari pre[target]
	seed, ok := pre[target]
	if !ok {
		return nil, fmt.Errorf("pre must contain target entry: %s", target)
	}

	// Local visited + steps + queue
	visited := map[string]bool{seed[0]: true, seed[1]: true}

	type state struct{ elem string }
	q := list.New()
	q.PushBack(state{seed[0]})
	q.PushBack(state{seed[1]})

	// BFS tanpa sentuh g.Recipes
	for q.Len() > 0 {
		curr := q.Remove(q.Front()).(state).elem

		if g.Tier(curr) == 0 {
			continue
		}

		// strict → fallback
		combos := g.RecipesFor(curr, true)
		if len(combos) == 0 {
			combos = g.RecipesFor(curr, false)
		}
		if len(combos) == 0 {
			continue
		}
		// seed pre[curr] sekali
		if _, exists := pre[curr]; !exists {
			pre[curr] = [2]string{combos[0][0], combos[0][1]}
		}
		for _, ing := range pre[curr] {
			if !visited[ing] {
				visited[ing] = true
				q.PushBack(state{ing})
			}
		}
	}

	tree := BuildTreeFromPre(g, target, pre)
	dur := time.Since(start)
	duration := fmt.Sprintf("%.3fms", float64(dur.Nanoseconds())/1e6)
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
	nodesVisited := countNodes(tree)
	return []SearchResult{{
		Recipe:       tree,
		NodesVisited: nodesVisited,
		Duration:     duration,
	}}, nil
}

// SearchBFSMultiple mencari hingga maxRecipes resep berbeda untuk `target` secara paralel.
func SearchBFSMultiple(
	g *utils.Graph,
	target string,
	maxRecipes int,
) ([]SearchResult, error) {
	// Ambil semua resep awal (strict lalu fallback)
	combos := g.RecipesFor(target, true)
	if len(combos) == 0 {
		combos = g.RecipesFor(target, false)
	}

	var (
		mu          sync.Mutex
		wg          sync.WaitGroup
		results     []SearchResult
		seen        = make(map[string]bool)
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer cancel()

	// Helper memanggil SearchBFSWithPre tanpa override metrik
	bfsWithPre := func(pre map[string][2]string) (*SearchResult, error) {
		res, err := SearchBFSWithPre(g, target, pre)
		if err != nil || len(res) == 0 {
			return nil, err
		}
		return &res[0], nil
	}

	for _, combo := range combos {
		combo := combo
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
			}

			pre := map[string][2]string{target: {combo[0], combo[1]}}
			r, err := bfsWithPre(pre)
			if err != nil {
				return
			}

			// deduplikasi berdasarkan struktur tree JSON
			sig, _ := json.Marshal(r.Recipe)
			key := string(sig)

			mu.Lock()
			defer mu.Unlock()
			if len(results) < maxRecipes && !seen[key] {
				seen[key] = true
				results = append(results, *r)
				if len(results) >= maxRecipes {
					cancel()
				}
			}
		}()
	}

	wg.Wait()
	if len(results) == 0 {
		return nil, fmt.Errorf("no complete recipes found for %q", target)
	}
	return results, nil
}
