package main

import (
	"fmt"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/rapito/go-shopify/shopify"
)

const (
	store  = "dummy-store-abcd"
	apiKey = "7c92182fcb56bd9232b3f3771f9b8ffd"
	pass   = "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"
)

// Create a new shopify object with your store
// domain, api key and password
var shop = shopify.New(store, apiKey, pass)

func main() {

	fetchAllProducts()
	fetchOneProduct(8008760820017)
}

func fetchOneProduct(id int64) {
	fmt.Println("[fetchOneProduct]")

	// Call any of the api CRUD methods
	endpoint := fmt.Sprintf("products/%v", id)
	result, _ := shop.Get(endpoint)

	fmt.Println("Result")
	fmt.Println(string(result))

	fmt.Println("===============")
}

func fetchAllProducts() {
	fmt.Println("[fetchAllProducts]")

	// Call any of the api CRUD methods
	result, _ := shop.Get("products")

	// Do what you want with the []byte response.
	// In this case we are using simplejson to handle it.
	jsonData, _ := simplejson.NewJson(result)

	fmt.Println("======+++**_+++=========")
	fmt.Println(jsonData["id"])

	products := jsonData.Get("products")
	product := products.GetIndex(0)

	title, _ := product.Get("title").String()

	fmt.Println("Full result: ")
	fmt.Println(string(result))

	fmt.Println("First Product title: ")
	fmt.Println(title)

	fmt.Println("===============")
}
