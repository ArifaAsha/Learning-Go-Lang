package main

import (
	"fmt" //formatted I/O

	"github.com/gocolly/colly" //scraping framework
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {

	c := colly.NewCollector(colly.AllowedDomains("www.amazon.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Link of the page:", r.URL)
	})

	var items []item //blank list or black slice

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.a-section.a-spacing-base", func(_ int, h *colly.HTMLElement) {
			item := item{
				Name:   h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal"),
				Price:  h.ChildText("span.a-price-whole"), //childext => child element
				ImgUrl: h.ChildAttr("img", "src"),
			}
			items = append(items, item)
		})
	})

	c.Visit("https://www.amazon.com/s?k=plushies")

	for _, item := range items {
		fmt.Printf("%s\n-----------\n", item)
	}

}
