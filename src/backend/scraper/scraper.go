package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// Fungsi untuk melakukan scraping dan menyimpan data dalam file JSON
func ScrapeElements(c *gin.Context) {
	// URL yang ingin di-scrape
	url := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)#List_of_elements" // Ganti dengan URL yang sesuai

	// Mengambil konten HTML dari URL
	res, err := http.Get(url)
	if err != nil {
		log.Println("Failed to fetch the page:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch the page",
		})
		return
	}
	defer res.Body.Close()

	// Mem-parsing HTML dengan Goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Failed to parse the page:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse the page",
		})
		return
	}

	// Menyimpan data elemen dan resep dalam map
	elementsRecipes := make(map[string][][2]string)

	// Menemukan tabel yang berisi elemen dan resep
	doc.Find("table.list-table.col-list.icon-hover tr").Each(func(i int, row *goquery.Selection) {
		// Mengabaikan header tabel
		if i > 0 {
			// Mengambil nama elemen
			element := row.Find("td").Eq(0).Text()
			element = strings.TrimSpace(element) // Menghapus spasi yang tidak perlu

			// Mengambil resep dari elemen tersebut
			var recipes [][2]string
			var prevRecipe string

			row.Find("td").Eq(1).Find("a").Each(func(j int, link *goquery.Selection) {
				recipe := link.Text()
				recipe = strings.TrimSpace(recipe) // Menghapus spasi yang tidak perlu
				// Hanya menambahkan pasangan bahan jika resep tidak kosong
				if recipe != "" {
					// Jika sudah ada bahan sebelumnya, buat pasangan
					if prevRecipe != "" {
						recipes = append(recipes, [2]string{prevRecipe, recipe})
						prevRecipe = "" // Reset bahan sebelumnya
					} else {
						// Jika belum ada pasangan, simpan bahan pertama
						prevRecipe = recipe
					}
				}
			})

			// Menyimpan pasangan elemen dan resep dalam map hanya jika pasangan lengkap
			if len(recipes) > 0 {
				elementsRecipes[element] = recipes
			}
		}
	})

	// Menyimpan data ke file JSON
	err = saveToJSON(elementsRecipes)
	if err != nil {
		log.Println("Failed to save data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save the data",
		})
		return
	}

	// Mengembalikan hasil dalam format JSON
	c.JSON(http.StatusOK, gin.H{
		"message": "Scraping successful, data saved to file.",
	})
}

// Fungsi untuk menyimpan data elemen dan resep dalam file JSON
func saveToJSON(data map[string][][2]string) error {
	// Menyusun file JSON
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Menyimpan ke dalam file menggunakan os.WriteFile
	err = os.WriteFile("scraped_recipes.json", file, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Data successfully saved to scraped_recipes.json")
	return nil
}
