package router

import (
//	"encoding/json"
	"net/http"
	"time"
	"tubes2/search"
	"tubes2/utils"

	"github.com/gin-gonic/gin"
)

// RequestBody adalah format JSON yang dikirim dari frontend
type RequestBody struct {
    Target      string `json:"target"`
    Algorithm   string `json:"algorithm"`   // "bfs", "dfs", "bidirectional"
    Mode        string `json:"mode"`        // "single", "multiple"
    MaxRecipes  int    `json:"max_recipes"` // untuk multiple
}

// ResponseBody hasil akhir yang dikembalikan ke frontend
type ResponseBody struct {
    Recipes   []search.TreeNode   `json:"recipes"`
    Steps     []int               `json:"steps"`
    Durations []string            `json:"durations"`
}

func SetupRouter(g *utils.Graph) *gin.Engine {
    router := gin.Default()

    router.POST("/search", func(c *gin.Context) {
        var req RequestBody
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }

        if req.Target == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Target is required"})
            return
        }

        var recipes []search.TreeNode
        var stepsList []int
        var durations []string

        if req.Mode == "single" {
            var recipe search.TreeNode
            var steps int
            var dur time.Duration

            switch req.Algorithm {
            case "bfs":
                recipe, steps, dur = search.BFS(g, req.Target)
            case "dfs":
                recipe, steps, dur = search.DFS(g, req.Target)
            case "bidirectional":
                recipe, steps, dur = search.BidirectionalSearch(g, req.Target)
            default:
                c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown algorithm"})
                return
            }

            recipes = []search.TreeNode{recipe}
            stepsList = []int{steps}
            durations = []string{dur.String()}

        } else if req.Mode == "multiple" {
            if req.MaxRecipes <= 0 {
                c.JSON(http.StatusBadRequest, gin.H{"error": "max_recipes must be positive"})
                return
            }

            var results []search.MultiResult
            switch req.Algorithm {
            case "bfs":
                results = search.BFSAll(g, req.Target, req.MaxRecipes)
            case "dfs":
                results = search.DFSAll(g, req.Target, req.MaxRecipes)
            case "bidirectional":
                results = search.BidirectionalAll(g, req.Target, req.MaxRecipes)
            default:
                c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown algorithm"})
                return
            }

            for _, r := range results {
                recipes = append(recipes, r.Recipe)
                stepsList = append(stepsList, r.Steps)
                durations = append(durations, r.Duration.String())
            }
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown mode"})
            return
        }

        c.JSON(http.StatusOK, ResponseBody{
            Recipes:   recipes,
            Steps:     stepsList,
            Durations: durations,
        })
    })

    return router
}
