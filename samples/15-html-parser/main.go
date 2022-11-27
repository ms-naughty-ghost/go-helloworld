// html解析
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://tinpangames.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".Home_card___LpL1").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("h2").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
	title := doc.Find("title").Text()
	fmt.Println(title)
}
func main() {
	ExampleScrape()
}
