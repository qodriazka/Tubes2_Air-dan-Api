package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// ElementData menampung tier dan daftar pasangan resep
type ElementData struct {
    Tier    int        `json:"tier"`
    Recipes [][2]string `json:"recipes"`
}

// ScrapeToFile menjalankan scraping dan menulis hasil ke file JSON
func ScrapeToFile(jsonPath string) error {
    url := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)#List_of_elements"
    res, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("failed to fetch page: %w", err)
    }
    defer res.Body.Close()
    if res.StatusCode != 200 {
        return fmt.Errorf("unexpected status code: %d", res.StatusCode)
    }

    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        return fmt.Errorf("failed to parse page: %w", err)
    }

    elementsData := make(map[string]ElementData)

    // Cari setiap section headline yang menandai tier
    doc.Find("span.mw-headline").Each(func(_ int, s *goquery.Selection) {
        title := strings.TrimSpace(s.Text())
        var tier int
        switch {
        case strings.EqualFold(title, "Starting elements"):
            tier = 0
        case strings.HasPrefix(title, "Tier"):
            // e.g. "Tier 1 elements"
            parts := strings.Fields(title)
            if len(parts) >= 2 {
                if n, err := strconv.Atoi(parts[1]); err == nil {
                    tier = n
                }
            }
        default:
            // bukan section recipe
            return
        }

        // Ambil tabel pertama setelah heading tersebut
        table := s.Parent().NextAll().Filter("table").First()
        if table.Length() == 0 {
            log.Printf("Warning: no table found for section %q", title)
            return
        }

        // Proses baris pada table
        table.Find("tr").Each(func(i int, row *goquery.Selection) {
            if i == 0 {
                return // skip header row
            }
            name := strings.TrimSpace(row.Find("td").Eq(0).Text())
            if name == "" {
                return
            }
            // parse resep
            var recipes [][2]string
            var prev string
            row.Find("td").Eq(1).Find("a").Each(func(_ int, link *goquery.Selection) {
                txt := strings.TrimSpace(link.Text())
                if txt == "" {
                    return
                }
                if prev == "" {
                    prev = txt
                } else {
                    recipes = append(recipes, [2]string{prev, txt})
                    prev = ""
                }
            })

            elementsData[name] = ElementData{
                Tier:    tier,
                Recipes: recipes,
            }
        })
    })

    // Tulis ke JSON
    dataBytes, err := json.MarshalIndent(elementsData, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal JSON: %w", err)
    }
    if err := os.WriteFile(jsonPath, dataBytes, 0644); err != nil {
        return fmt.Errorf("failed to write JSON file: %w", err)
    }
    return nil
}

// ScrapeElements adalah HTTP handler untuk scraping via API
func ScrapeElements(c *gin.Context) {
    if err := ScrapeToFile("scraped_recipes.json"); err != nil {
        log.Println("Scraping failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Scraping successful, data saved to scraped_recipes.json"})
}
