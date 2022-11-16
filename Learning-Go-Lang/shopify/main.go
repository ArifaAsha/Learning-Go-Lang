package main

// package main

// import (
// 	"fmt"

// 	"github.com/arduino/go-shopify/shopify"
// )

// func main() {
// 	shop := shopify.NewClient("dummy-3378", "523f4d4460cd849e2dc16bbb9bf2a24b", "https://dummy-3378.myshopify.com/admin/api/2022-10/products.json")
// 	shop.LoadProducts()
// 	// fmt.Println(shop)
// 	fmt.Printf("%v\n", shop.Products)
// }

import (
	shopify "github.com/boourns/go_shopify"
)

func main() {
	api = shopify.API{
		URI:    "https://dummy-3378.myshopify.com/admin/",
		Token:  "3649e565061a54216c506a876cbbc6db",
		Secret: "523f4d4460cd849e2dc16bbb9bf2a24b",
	}

	products := api.Products()
	// or
	product := api.Product(12345)
}
