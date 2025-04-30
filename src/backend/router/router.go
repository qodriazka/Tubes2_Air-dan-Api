package router

import (
	"tubes2/scraper" // Mengimpor package scraper yang benar

	"github.com/gin-gonic/gin"
)

// Fungsi untuk mengatur routing
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routing untuk endpoint scraping
	r.GET("/scraper", scraper.ScrapeElements) // Memanggil fungsi ScrapeElements dari scraper

	return r
}
