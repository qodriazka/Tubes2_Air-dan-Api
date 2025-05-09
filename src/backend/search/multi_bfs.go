package search

import (
    "sync"
    "tubes2/utils"
    "time"
)

// MultiResult menyimpan hasil satu pencarian
type MultiResult struct {
    Recipe   TreeNode
    Steps    int
    Duration time.Duration
}

func (t TreeNode) Signature() string {
    if len(t.Combines) == 0 {
        return t.Name
    }
    leftSig  := t.Combines[0].Signature()
    rightSig := t.Combines[1].Signature()
    return t.Name + "(" + leftSig + "," + rightSig + ")"
}

// BFSAll menjalankan persis maxRecipes BFS secara parallel, 
// dan hanya mengumpulkan resep unik hingga maxRecipes.
func BFSAll(
    g *utils.Graph,
    target string,
    maxRecipes int,
) []MultiResult {
    var wg sync.WaitGroup
    resultsCh := make(chan MultiResult, maxRecipes)

    // set untuk memastikan unik
    seen := make(map[string]struct{})
    var mu sync.Mutex

    // spawn satu goroutine per recipe diminta
    for i := 0; i < maxRecipes; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()

            tree, steps, dur := BFS(g, target)

			//if tree == nil {
			//	return
			//}
            // compute signature untuk unik
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

    // saat semua goroutine selesai, tutup channel
    go func() {
        wg.Wait()
        close(resultsCh)
    }()

    // kumpulkan hingga habis atau sudah dapat maxRecipes unik
    var results []MultiResult
    for res := range resultsCh {
        results = append(results, res)
        if len(results) >= maxRecipes {
            break
        }
    }
    return results
}
