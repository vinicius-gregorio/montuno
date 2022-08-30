package scrapper

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type House struct {
	Title   string
	Address string
	Price   string
	Area    string
}

func StartScrappingHouses() {
	url := "https://www.vivareal.com.br/aluguel/sp/santos/"

	houses := []House{}
	for i := 0; i < 5; i++ {
		if i != 0 {
			url = "https://www.vivareal.com.br/aluguel/sp/santos/?pagina=" + strconv.FormatInt(int64(i), 10)
		}
		c := colly.NewCollector()

		c.OnHTML(".property-card__container", func(e *colly.HTMLElement) {
			house := House{}
			house.Title = e.ChildText(".js-card-title")
			house.Address = e.ChildText(".property-card__address")
			house.Area = e.ChildText(".js-property-card-detail-area") + "m²"
			house.Price = e.ChildText(".property-card__price")
			houses = append(houses, house)
			fmt.Println(house)
		})
		c.OnResponse(func(r *colly.Response) {
			fmt.Println(r.StatusCode)
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		c.Visit(url)
	}

	writeHousesToJson(houses)

}

func writeHousesToJson(houses []House) {
	file, _ := json.MarshalIndent(houses, "", "")
	os.WriteFile("outputs\\output-houses.json", file, 0644)
}
