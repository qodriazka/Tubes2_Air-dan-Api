package router

import (
	"net/http"
	"strconv"
	"time"
	"tubes2/scraper" // Mengimpor package scraper yang benar
	"tubes2/search"
	"tubes2/utils"
	"github.com/gin-gonic/gin"
)

// Fungsi untuk mengatur routing
func SetupRouter(g *utils.Graph) *gin.Engine {
	r := gin.Default()

	// Routing untuk endpoint scraping
	r.GET("/scraper", scraper.ScrapeElements) // Memanggil fungsi ScrapeElements dari scraper

	// Endpoint list elemen
	r.GET("/elements", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"elements": g.Nodes()})
	})
	
	// Endpoint search
	r.GET("/search", func(c *gin.Context) {
		start := c.DefaultQuery("start","Water")
		target := c.Query("target")
		algo   := c.DefaultQuery("algo", "bfs") // bfs atau dfs
		mode := c.DefaultQuery("mode", "multiple")
		maxStr := c.DefaultQuery("max", "3")

		if target == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "`target` is required"})
			return
		}

		if _, ok := g.Adj[start]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "`start` element not found"})
			return
		}

		if _, ok := g.Adj[target]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "`target` element not found"})
			return
		}

		if mode == "multiple"{
			max, err := strconv.Atoi(maxStr)
			if err != nil || max <= 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "`max` must be a positive integer"})
				return
			}
			paths := search.FindMultiplePaths(g, start, target, algo, max)
			c.JSON(http.StatusOK, gin.H{
				"paths":	paths,
				"count":	len(paths),
			})	
			return
		}

		var (
		  path    []string
		  visited int
		  dur     time.Duration
		)
		if algo == "dfs" {
		  path, visited, dur = search.DFS(g, start, target)
		} else {
		  path, visited, dur = search.BFS(g, start, target)
		}
	
		c.JSON(http.StatusOK, gin.H{
		  "path":    path,
		  "visited": visited,
		  "time_ms": dur.Milliseconds(),
		})
	  })
	return r
}
