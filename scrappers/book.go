package scrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func StartScrapingBooks() {
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://books.toscrape.com/")
}
