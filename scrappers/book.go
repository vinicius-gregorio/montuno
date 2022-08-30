package scrapper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Book struct {
	Title string
	Price string
}

func StartScrapingBooks() {
	const url = "https://books.toscrape.com/"
	const allowedDomains = "books.toscrape.com"
	fileCSV, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	books := []Book{}

	defer fileCSV.Close()

	writerCSV := csv.NewWriter(fileCSV)
	defer writerCSV.Flush()

	headers := []string{"Title", "Price"}
	writerCSV.Write(headers)

	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomains),
	)

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		book := Book{}
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.Price = e.ChildText(".price_color")
		fmt.Println(book.Title, book.Price)
	})

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		book := Book{}
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.Price = e.ChildText(".price_color")
		row := []string{book.Title, book.Price}
		writerCSV.Write(row)
		books = append(books, book)

	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	writeBooksToJson(books)
}

func writeBooksToJson(books []Book) {
	file, _ := json.MarshalIndent(books, "", "")
	ioutil.WriteFile("output.json", file, 0644)
}
