package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeElements function to scrape all elements and their recipes
func ScrapeElements() ([]map[string]interface{}, error) {
	// URL halaman wiki Little Alchemy 2
	url := "https://little-alchemy.fandom.com/wiki/Elements_(Little_Alchemy_2)"

	// Lakukan request ke halaman web
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page: %v", err)
	}
	defer res.Body.Close()

	// Pastikan response OK (status 200)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch page, status code: %d", res.StatusCode)
	}

	// Parse HTML dari response menggunakan goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	// Slice untuk menyimpan data elemen dan resep
	var elements []map[string]interface{}

	// Loop untuk mencari elemen-elemen pada halaman
	doc.Find(".wikitable tbody tr").Each(func(i int, s *goquery.Selection) {
		// Ambil kolom pertama untuk nama elemen
		element := s.Find("td").Eq(0).Text()
		element = strings.TrimSpace(element)

		// Jika elemen kosong, skip
		if element == "" {
			return
		}

		// Ambil resep di kolom kedua (jika ada)
		recipes := s.Find("td").Eq(1).Text()
		recipes = strings.TrimSpace(recipes)

		// Jika tidak ada resep, skip
		if recipes == "" {
			return
		}

		// Simpan data dalam bentuk map
		elements = append(elements, map[string]interface{}{
			"element": element,
			"recipes": recipes,
		})
	})

	return elements, nil
}

func main() {
	// Memanggil fungsi ScrapeElements
	elements, err := ScrapeElements()
	if err != nil {
		log.Fatal(err)
	}

	// Menampilkan hasil scraping
	for _, element := range elements {
		fmt.Printf("Element: %s\n", element["element"])
		fmt.Printf("Recipes: %s\n\n", element["recipes"])
	}
}
