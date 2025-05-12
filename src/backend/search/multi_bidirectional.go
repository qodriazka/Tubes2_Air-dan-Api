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

// SearchBidirectionalWithSeed is a variant of two-way BFS seeded with a fixed initial combo for target.
func SearchBidirectionalWithSeed(
	g *utils.Graph,
	target string,
	fixed [2]string,
) (*SearchResult, error) {
	start := time.Now()

	// Forward frontier seeded with fixed ingredients
	fQueue := list.New()
	fVis := make(map[string]bool)
	fPre := make(map[string][2]string)
	// seed pre[target]
	fPre[target] = fixed
	// seed queue and visited
	for _, ing := range fixed {
		fQueue.PushBack(ing)
		fVis[ing] = true
	}

	// Backward frontier from all base elements
	bQueue := list.New()
	bVis := make(map[string]bool)
	bPre := make(map[string][2]string)
	for name, elem := range g.Elements {
		if elem.Tier == 0 {
			bQueue.PushBack(name)
			bVis[name] = true
		}
	}

	//steps := 0
	var meetNode string
	var meetPair [2]string

	// expand interleaved
	for fQueue.Len() > 0 && bQueue.Len() > 0 {
		// forward step
		//steps++
		meetNode, meetPair = expandForward(g, fQueue, fVis, fPre, bVis)
		if meetNode != "" {
			break
		}
		// backward step
		//steps++
		meetNode, meetPair = expandBackward(g, bQueue, bVis, bPre, fVis)
		if meetNode != "" {
			break
		}
	}

	if meetNode == "" {
		return nil, fmt.Errorf("no path found for %q with seed %v", target, fixed)
	}

	// combine pre maps
	jointPre := make(map[string][2]string)
	for k, v := range fPre {
		jointPre[k] = v
	}
	for k, v := range bPre {
		jointPre[k] = v
	}
	jointPre[meetNode] = meetPair

	// reconstruct
	tree := BuildTreeFromPreSafe(g, target, jointPre)
	duration := fmt.Sprintf("%.3fms", float64(time.Since(start).Nanoseconds())/1e6)
	// compute nodesVisited as total nodes in recipe tree
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
	return &SearchResult{Recipe: tree, NodesVisited: nodesVisited, Duration: duration}, nil
}

// SearchBidirectionalMultiple finds up to maxRecipes distinct recipes using bidirectional search per initial combo.
func SearchBidirectionalMultiple(
	g *utils.Graph,
	target string,
	maxRecipes int,
) ([]SearchResult, error) {
	combos := g.RecipesFor(target, false)
	if len(combos) == 0 {
		return nil, fmt.Errorf("no initial combos for %q", target)
	}

	var (
		mu          sync.Mutex
		wg          sync.WaitGroup
		results     []SearchResult
		seen        = make(map[string]bool)
		ctx, cancel = context.WithCancel(context.Background())
	)
	defer cancel()

	for _, combo := range combos {
		if len(combo) != 2 {
			continue
		}
		fixed := [2]string{combo[0], combo[1]}
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
			}
			res, err := SearchBidirectionalWithSeed(g, target, fixed)
			if err != nil {
				return
			}
			sig, _ := json.Marshal(res.Recipe)
			key := string(sig)

			mu.Lock()
			defer mu.Unlock()
			if len(results) < maxRecipes && !seen[key] {
				seen[key] = true
				results = append(results, *res)
				if len(results) >= maxRecipes {
					cancel()
				}
			}
		}()
	}
	wg.Wait()
	if len(results) == 0 {
		return nil, fmt.Errorf("no recipes found for %q", target)
	}
	return results, nil
}
