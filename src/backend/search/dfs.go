package search

import (
    "time"
    "tubes2/utils"
)

// DFS melakukan traversal depth-first dan stop di node lengkap pertama
func DFS(g *utils.Graph, target string) (TreeNode, int, time.Duration) {
    tid := g.IDs[target]
    n := len(g.AdjInt)

    visited := make([]bool, n)
    steps := 0
    t0 := time.Now()

    var dfs func(id int) (TreeNode, bool)
    dfs = func(id int) (TreeNode, bool) {
        if visited[id] {
            return TreeNode{}, false
        }
        visited[id] = true
        steps++

        name := g.Names[id]
        if primitives[name] {
            return TreeNode{Name: name}, false
        }

        children := g.AdjInt[id]
        if len(children) == 2 {
            leftName := g.Names[children[0]]
            rightName := g.Names[children[1]]
            if primitives[leftName] && primitives[rightName] {
                left := TreeNode{Name: leftName}
                right := TreeNode{Name: rightName}
                return TreeNode{
                    Name:     name,
                    Combines: []TreeNode{left, right},
                }, true
            }
        }

        for _, v := range children {
            if _, ok := dfs(v); ok {
                return buildTreeNode(id, g), true
            }
        }

        return TreeNode{}, false
    }

    result, _ := dfs(tid)
    return result, steps, time.Since(t0)
}