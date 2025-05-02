package search

import (
    "fmt"
    "sync"
    "tubes2/utils"
)

func pathKey(path []string) string {
    return fmt.Sprintf("%q", path)
}

func FindMultiplePaths(g *utils.Graph, start, target, algo string, max int) [][]string {
    if algo == "bfs" {
        return FindMultipleBFS(g, start, target, max)
    }
    return FindMultipleDFS(g, start, target, max)
}

func FindMultiplePathsInt(g *utils.Graph, start, target, algo string, max int) [][]string {
    if algo == "bfs" {
        return FindMultipleBFSInt(g, start, target, max)
    }
    return FindMultipleDFSInt(g, start, target, max)
}

func FindMultipleDFS(g *utils.Graph, start, target string, max int) [][]string {
	var mu sync.Mutex
	results := make([][]string, 0, max)
	seen := make(map[string]bool)

	var wg sync.WaitGroup

	var dfsAll func(path []string, visited map[string]bool)
	dfsAll = func(path []string, visited map[string]bool){
		if len(results) >= max {
			return
		}

		curr := path[len(path)-1]
		if curr == target {
			key := pathKey(path)
			mu.Lock()
			if !seen[key] && len(results) < max {
				results = append(results, append([]string{}, path...))
				seen[key] = true
			}
			mu.Unlock()
			return
		}

		for _, nbr := range g.Adj[curr]{
			if !visited[nbr] {
				visited[nbr] = true
				dfsAll(append(path, nbr), visited)
				visited[nbr] = false
				if len(results) >= max {
					return
				}
			}
		}
	}

	for _, first := range g.Adj[start]{
		wg.Add(1)
		go func(nbr string){
			defer wg.Done()
			visited := map[string]bool{start: true, nbr: true}
			dfsAll([]string{start, nbr}, visited)
		}(first)
	}

	wg.Wait()
	return results
}

func FindMultipleBFS(g *utils.Graph, start, target string, max int) [][]string {
    var mu sync.Mutex
    results := make([][]string, 0, max)
    seen := make(map[string]bool)

    queue := [][]string{{start}}
    for len(queue) > 0 && len(results) < max {
        path := queue[0]
        queue := queue[1:]
        last := path[len(path)-1]

        for _, nbr := range g.Adj[last]{
            if contains(path, nbr) {
                continue
            }
            
            newPath := append([]string{}, path...)
            newPath = append(newPath, nbr)

            if nbr == target {
                key := fmt.Sprintf("%q", newPath)
                mu.Lock()
                if !seen[key] && len(results) < max {
                    results = append(results, newPath)
                    seen[key] = true
                }
                mu.Unlock()
                if len(results) >= max {
                    break
                } else {
                    queue = append(queue, newPath)
                }
            }
        }

    }
    return results
}

func contains(s []string, v string) bool {
    for _, x := range s {
        if x == v {
            return true
        }
    }

    return false
}

// findMultipleBFSInt: layer‐by‐layer exploration hingga max paths
func FindMultipleBFSInt(g *utils.Graph, startName, targetName string, max int) [][]string {
    startID := g.IDs[startName]
    targetID := g.IDs[targetName]

    var mu sync.Mutex
    results := make([][]string, 0, max)
    seen := make(map[string]bool)

    // queue of ID paths
    queue := [][]int{{startID}}

    for len(queue) > 0 && len(results) < max {
        pathIDs := queue[0]
        queue = queue[1:]
        last := pathIDs[len(pathIDs)-1]

        for _, nbr := range g.AdjInt[last] {
            // avoid cycle
            if containsID(pathIDs, nbr) {
                continue
            }
            newPath := append([]int{}, pathIDs...)
            newPath = append(newPath, nbr)

            if nbr == targetID {
                key := pathKeyID(newPath, g.Names)
                mu.Lock()
                if !seen[key] && len(results) < max {
                    // convert IDs → names
                    results = append(results, idsToNames(newPath, g.Names))
                    seen[key] = true
                }
                mu.Unlock()
                if len(results) >= max {
                    break
                }
            } else {
                queue = append(queue, newPath)
            }
        }
    }
    return results
}

// findMultipleDFSInt: backtracking paralel per neighbor pertama
func FindMultipleDFSInt(g *utils.Graph, startName, targetName string, max int) [][]string {
    startID := g.IDs[startName]
    targetID := g.IDs[targetName]

    var mu sync.Mutex
    results := make([][]string, 0, max)
    seen := make(map[string]bool)
    var wg sync.WaitGroup

    var dfsAll func(pathIDs []int, visited map[int]bool)
    dfsAll = func(pathIDs []int, visited map[int]bool) {
        if len(results) >= max {
            return
        }
        curr := pathIDs[len(pathIDs)-1]
        if curr == targetID {
            key := pathKeyID(pathIDs, g.Names)
            mu.Lock()
            if !seen[key] && len(results) < max {
                results = append(results, idsToNames(pathIDs, g.Names))
                seen[key] = true
            }
            mu.Unlock()
            return
        }
        for _, nbr := range g.AdjInt[curr] {
            if !visited[nbr] {
                visited[nbr] = true
                dfsAll(append(pathIDs, nbr), visited)
                visited[nbr] = false
                if len(results) >= max {
                    return
                }
            }
        }
    }

    // paralel untuk tiap neighbor pertama
    for _, nbr := range g.AdjInt[startID] {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            visited := map[int]bool{startID: true, n: true}
            dfsAll([]int{startID, n}, visited)
        }(nbr)
    }
    wg.Wait()
    return results
}

// containsID cek apakah slice of int mengandung value
func containsID(s []int, v int) bool {
    for _, x := range s {
        if x == v {
            return true
        }
    }
    return false
}

// pathKeyID untuk deduplikasi: serialisasi path of IDs menjadi string
func pathKeyID(pathIDs []int, names []string) string {
    // gunakan nama agar unik dan mudah dibaca
    strs := idsToNames(pathIDs, names)
    return fmt.Sprintf("%q", strs)
}

// idsToNames convert []int → []string
func idsToNames(pathIDs []int, names []string) []string {
    out := make([]string, len(pathIDs))
    for i, id := range pathIDs {
        out[i] = names[id]
    }
    return out
}