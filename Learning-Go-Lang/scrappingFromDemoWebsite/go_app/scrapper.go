package main

// https://www.youtube.com/watch?v=NU4OlJVj1gs

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"), //only the secific domain information
	) //instance for a web scrapper

	// c.OnHTML("div[itemprop=itemListElement]", func(h *colly.HTMLElement) { //find itemprop and print all the text
	// 	fmt.Println(h.Text)
	// })

	// c.OnHTML("div[itemprop=itemListElement]", func(h *colly.HTMLElement) { //find itemprop and print all the text
	// 	fmt.Println(h.ChildText("h2.product-title")) // h2 tag and class => product title
	// }) //title of all of the products on the page

	// c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) { //find itemprop and print all the text
	// 	fmt.Println(h.ChildText("h2.product-title")) // h2 tag and class => product title
	// }) //without space at the last

	var items []item //blank list or black slice

	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) { //find itemprop and print all the text
		item := item{
			Name:   h.ChildText("h2.product-title"),
			Price:  h.ChildText("div.sale-price"), //childext => child element
			ImgUrl: h.ChildAttr("img", "src"),
		}
		// h.ChildText("h2.product-title") // h2 tag and class => product title
		// fmt.Println(item)
		items = append(items, item)
	})

	c.Visit("http://j2store.net/demo/index.php/shop")
	fmt.Println(items)
}
