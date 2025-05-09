package search

import (
    "sync"
    "tubes2/utils"
)

func DFSAll(
    g *utils.Graph,
    target string,
    maxRecipes int,
) []MultiResult {
    var wg sync.WaitGroup
    resultsCh := make(chan MultiResult, maxRecipes)

    seen := make(map[string]struct{})
    var mu sync.Mutex

    for i := 0; i < maxRecipes; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()

            tree, steps, dur := DFS(g, target)
            sig := tree.Signature()

            mu.Lock()
            defer mu.Unlock()
            if _, exists := seen[sig]; !exists {
                seen[sig] = struct{}{}
                resultsCh <- MultiResult{
                    Recipe:   tree,
                    Steps:    steps,
                    Duration: dur,
                }
            }
        }()
    }

    go func() {
        wg.Wait()
        close(resultsCh)
    }()

    var results []MultiResult
    for res := range resultsCh {
        results = append(results, res)
        if len(results) >= maxRecipes {
            break
        }
    }
    return results
}
