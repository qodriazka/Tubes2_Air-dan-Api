package router

import (
	"net/http"
	"tubes2/search"
	"tubes2/utils"

	"github.com/gin-contrib/cors" // Menambahkan impor untuk CORS
	"github.com/gin-gonic/gin"
)

// RequestBody adalah format JSON yang dikirim dari frontend
type RequestBody struct {
	Target     string `json:"target"`
	Algorithm  string `json:"algorithm"`   // "bfs", "dfs"
	Mode       string `json:"mode"`        // "single", "multiple"
	MaxRecipes int    `json:"max_recipes"` // untuk multiple
}

// ResponseBody hasil akhir yang dikembalikan ke frontend
// Disusun sebagai array SearchResult JSON langsung

// SetupRouter membuat router HTTP dengan satu endpoint /search
func SetupRouter(g *utils.Graph) *gin.Engine {
	r := gin.Default()

	// Menambahkan middleware CORS
	r.Use(cors.Default()) // Ini memungkinkan frontend di localhost:3000 untuk mengakses backend di localhost:8080

	r.POST("/search", func(c *gin.Context) {
		var req RequestBody
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// validasi target ada di graph
		if req.Target == "" || g.Tier(req.Target) < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown or missing target element"})
			return
		}

		// validasi mode & algorithm
		if req.Mode != "single" && req.Mode != "multiple" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mode"})
			return
		}
		if req.Algorithm != "bfs" && req.Algorithm != "dfs" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid algorithm"})
			return
		}

		// dispatch
		var (
			results []search.SearchResult
			err     error
		)

		if req.Mode == "single" {
			switch req.Algorithm {
			case "bfs":
				results, err = search.SearchBFS(g, req.Target)
			case "dfs":
				results, err = search.SearchDFS(g, req.Target)
			}
		} else {
			if req.MaxRecipes <= 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "max_recipes must be positive"})
				return
			}
			if req.Algorithm == "bfs" {
				results, err = search.SearchBFSMultiple(g, req.Target, req.MaxRecipes)
			} else {
				results, err = search.SearchDFSMultiple(g, req.Target, req.MaxRecipes)
			}
		}

		if err != nil {
			// jika error pencarian
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, results)
	})

	return r
}
