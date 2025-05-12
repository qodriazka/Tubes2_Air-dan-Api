package search

import (
    "container/list"
    "fmt"
    "time"
    "tubes2/utils"
)

// SearchBidirectional melakukan two-way BFS pada element graph
func SearchBidirectional(g *utils.Graph, target string) ([]SearchResult, error) {
    start := time.Now()

    // Forward frontier dari target
    fQueue := list.New()
    fQueue.PushBack(target)
    fVis := map[string]bool{target: true}
    fPre := make(map[string][2]string)

    // Backward frontier dari semua elemen dasar (tier 0)
    bQueue := list.New()
    bVis := make(map[string]bool)
    bPre := make(map[string][2]string)
    for name, elem := range g.Elements {
        if elem.Tier == 0 {
            bQueue.PushBack(name)
            bVis[name] = true
        }
    }

    steps := 0
    var meetNode string
    var meetPair [2]string

    // Expand interleaved until meet
    for fQueue.Len() > 0 && bQueue.Len() > 0 {
        // Forward layer
        steps++
        meetNode, meetPair = expandForward(g, fQueue, fVis, fPre, bVis)
        if meetNode != "" {
            break
        }
        // Backward layer
        steps++
        meetNode, meetPair = expandBackward(g, bQueue, bVis, bPre, fVis)
        if meetNode != "" {
            break
        }
    }

    if meetNode == "" {
        return nil, fmt.Errorf("no path found for %q", target)
    }

    // Gabungkan pre maps
    jointPre := make(map[string][2]string)
    for k, v := range fPre { jointPre[k] = v }
    for k, v := range bPre { jointPre[k] = v }
    jointPre[meetNode] = meetPair

    // Rekonstruksi pohon resep
    tree := BuildTreeFromPreSafe(g, target, jointPre)
    duration := fmt.Sprintf("%.3fms", float64(time.Since(start).Nanoseconds())/1e6)

    return []SearchResult{{Recipe: tree, NodesVisited: steps, Duration: duration}}, nil
}

// expandForward satu layer BFS maju
func expandForward(
    g *utils.Graph,
    q *list.List,
    vis map[string]bool,
    pre map[string][2]string,
    otherVis map[string]bool,
) (string, [2]string) {
    size := q.Len()
    for i := 0; i < size; i++ {
        curr := q.Remove(q.Front()).(string)
        if g.Tier(curr) == 0 {
            continue
        }
        combos := g.RecipesFor(curr, true)
        if len(combos) == 0 {
            combos = g.RecipesFor(curr, false)
        }
        for _, c := range combos {
            left, right := c[0], c[1]
            if _, ok := pre[curr]; !ok {
                pre[curr] = [2]string{left, right}
            }
            for _, ing := range []string{left, right} {
                if !vis[ing] {
                    vis[ing] = true
                    q.PushBack(ing)
                }
                if otherVis[ing] {
                    return ing, pre[curr]
                }
            }
        }
    }
    return "", [2]string{}
}

// expandBackward satu layer BFS mundur
func expandBackward(
    g *utils.Graph,
    q *list.List,
    vis map[string]bool,
    pre map[string][2]string,
    otherVis map[string]bool,
) (string, [2]string) {
    size := q.Len()
    for i := 0; i < size; i++ {
        child := q.Remove(q.Front()).(string)
        // scan semua resep untuk menemukan parent yang menggunakan child
        for parent, combos := range g.Recipes {
            for _, c := range combos {
                if c[0] == child || c[1] == child {
                    if _, ok := pre[parent]; !ok {
                        pre[parent] = [2]string{c[0], c[1]}
                    }
                    if !vis[parent] {
                        vis[parent] = true
                        q.PushBack(parent)
                    }
                    if otherVis[parent] {
                        return parent, pre[parent]
                    }
                }
            }
        }
    }
    return "", [2]string{}
}
func BuildTreeFromPreSafe(g *utils.Graph, target string, pre map[string][2]string) *Node {
    root := &Node{Name: target}
    visited := make(map[string]bool)

    var build func(n *Node)
    build = func(n *Node) {
        if visited[n.Name] {
            return
        }
        visited[n.Name] = true

        pair, ok := pre[n.Name]
        if !ok {
            // leaf/base, stop
            return
        }
        left := &Node{Name: pair[0]}
        right := &Node{Name: pair[1]}
        n.Combines = []*Node{left, right}

        build(left)
        build(right)
    }

    build(root)
    return root
}