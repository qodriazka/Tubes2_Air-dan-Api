package router

import (
    "net/http"
    "strconv"
    "time"

    "tubes2/scraper"
    "tubes2/search"
    "tubes2/utils"

    "github.com/gin-gonic/gin"
)

func SetupRouter(g *utils.Graph, recipes map[string][][2]string) *gin.Engine {
    r := gin.Default()

    // Endpoint untuk scraping
    r.GET("/scraper", scraper.ScrapeElements)

    // Endpoint untuk daftar semua elemen
    r.GET("/elements", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"elements": g.Nodes()})
    })

    // Endpoint pencarian
    r.GET("/search", func(c *gin.Context) {
        target := c.Query("target")
        algo   := c.DefaultQuery("algo", "bfs")
        mode   := c.DefaultQuery("mode", "single")
        maxStr := c.DefaultQuery("max", "3")

        if target == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "`target` is required"})
            return
        }
        if _, ok := g.IDs[target]; !ok {
            c.JSON(http.StatusBadRequest, gin.H{"error": "`target` element not found"})
            return
        }

        if mode == "multiple" {
            max, err := strconv.Atoi(maxStr)
            if err != nil || max <= 0 {
                c.JSON(http.StatusBadRequest, gin.H{"error": "`max` must be a positive integer"})
                return
            }
            paths := search.FindMultiplePaths(g, target, algo, max)
            c.JSON(http.StatusOK, gin.H{
                "paths": paths,
                "count": len(paths),
            })
            return
        }

        // mode == "single"
        var (
            path    []string
            visited int
            dur     time.Duration
        )
        switch algo {
        case "dfs":
            path, visited, dur = search.DFS(g, target)
        default:
            path, visited, dur = search.BFS(g, target)
        }

        c.JSON(http.StatusOK, gin.H{
            "path":    path,
            "visited": visited,
            "time_ms": dur.Milliseconds(),
        })
    })

    // Endpoint visualisasi pohon resep
    r.GET("/tree", func(c *gin.Context) {
        target := c.Query("target")
        if target == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "`target` is required"})
            return
        }
        tree := search.BuildRecipeTree(target, recipes)
        c.JSON(http.StatusOK, tree)
    })

    return r
}
