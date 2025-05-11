package search

import (
    "context"
    "encoding/json"
    "fmt"
    "sync"
    "time"
    "tubes2/utils"
)

func SearchDFSWithPre(
    g *utils.Graph,
    target string,
    pre map[string][2]string,
) ([]SearchResult, error) {
    start := time.Now()
    // track recursion stack to avoid cycles, not a global visited
    inPath := make(map[string]bool)
    var steps int

    // exactly the same DFS logic as SearchDFS, but seeded with pre[target]
    var dfs func(curr string) bool
    dfs = func(curr string) bool {
        steps++
        // base = tier 0
        if g.Tier(curr) == 0 {
            return true
        }
        if inPath[curr] {
            return false
        }
        inPath[curr] = true

        // try strict, fallback
        combos := g.RecipesFor(curr, true)
        if len(combos) == 0 {
            combos = g.RecipesFor(curr, false)
        }

        for _, combo := range combos {
            left, right := combo[0], combo[1]
            if dfs(left) && dfs(right) {
                // record only on successful branch
                if _, ok := pre[curr]; !ok {
                    pre[curr] = [2]string{left, right}
                }
                inPath[curr] = false
                return true
            }
        }

        inPath[curr] = false
        return false
    }

    // if caller pre-seeded pre[target], we honor that first
    if initial, ok := pre[target]; ok {
        // force the first step
        if !dfs(initial[0]) || !dfs(initial[1]) {
            return nil, fmt.Errorf("seeded recipe %v for %q failed to expand", initial, target)
        }
    }

    // now search the rest of the tree
    if !dfs(target) {
        return nil, fmt.Errorf("no recipe path found for %q", target)
    }

    // rebuild the tree
    tree := BuildTreeFromPre(g, target, pre)

    // format duration
    dur := time.Since(start)
    duration := fmt.Sprintf("%.3fms", float64(dur.Nanoseconds())/1e6)

    return []SearchResult{{
        Recipe:       tree,
        NodesVisited: steps,
        Duration:     duration,
    }}, nil
}


// SearchDFSMultiple finds up to maxRecipes distinct recipes for target
// by spawning DFS-per-initial-combo and collecting unique results.
func SearchDFSMultiple(
    g *utils.Graph,
    target string,
    maxRecipes int,
) ([]SearchResult, error) {
    combos := g.RecipesFor(target, false)
    if len(combos) == 0 {
        return nil, fmt.Errorf("no recipes to build %q", target)
    }

    var (
        mu       sync.Mutex
        wg       sync.WaitGroup
        results  []SearchResult
        seen     = make(map[string]bool)
        ctx, cancel = context.WithCancel(context.Background())
    )
    defer cancel()

    start := time.Now()
    dfsWithPre := func(initial [2]string) (*SearchResult, error) {
        pre := map[string][2]string{target: initial}
        res, err := SearchDFSWithPre(g, target, pre)
        if err != nil || len(res) == 0 {
            return nil, err
        }
        r := &res[0]
        r.Duration = time.Since(start).String()
        return r, nil
    }

    for _, combo := range combos {
        combo := combo
        if len(combo) != 2 {
            continue
        }
        wg.Add(1)
        go func() {
            defer wg.Done()
            select {
            case <-ctx.Done():
                return
            default:
            }
            r, err := dfsWithPre([2]string{combo[0], combo[1]})
            if err != nil {
                return
            }
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
