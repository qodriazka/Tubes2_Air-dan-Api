package search

import (
    "time"
    "tubes2/utils"
)

// BFS melakukan traversal level-order dan stop di node lengkap pertama
func BFS(g *utils.Graph, target string) (TreeNode, int, time.Duration) {
    tid := g.IDs[target]
    n   := len(g.AdjInt)

    // visited flags per run
    visited := make([]bool, n)
    // standard FIFO queue of IDs
    queue := []int{tid}
    visited[tid] = true

    steps := 0
    t0 := time.Now()

    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        steps++

        children := g.AdjInt[u]
        // stop condition: exactly 2 children, both primitives
        if len(children) == 2 {
            leftName  := g.Names[children[0]]
            rightName := g.Names[children[1]]
            if primitives[leftName] && primitives[rightName] {
                // reconstruct nested subtree and return
                result := buildTreeNode(u, g)
                return result, steps, time.Since(t0)
            }
        }

        // enqueue all unvisited neighbours
        for _, v := range children {
            if !visited[v] {
                visited[v] = true
                queue        = append(queue, v)
            }
        }
    }

    // jika tidak ditemukan sama sekali, return empty TreeNode
    return TreeNode{}, steps, time.Since(t0)
}

// buildTreeNode membangun subtree dari ID ke primitives, nested
func buildTreeNode(id int, g *utils.Graph) TreeNode {
    name := g.Names[id]
    if primitives[name] {
        return TreeNode{Name: name}
    }
    children := g.AdjInt[id]
    // asumsikan children length 2 untuk non-primitive dalam konteks ini
    leftNode  := buildTreeNode(children[0], g)
    rightNode := buildTreeNode(children[1], g)
    return TreeNode{
        Name:     name,
        Combines: []TreeNode{leftNode, rightNode},
    }
}
