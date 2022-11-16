package main

// https://github.com/rapito/go-shopify

import (
	"fmt"

	"github.com/rapito/go-shopify/shopify"
)

func main() {
	//1st= store name, 2nd = API key, 3rd = Admin API access token
	shop := shopify.New("dummy-store-abcd", "7c92182fcb56bd9232b3f3771f9b8ffd", "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")
	result, _ := shop.Get("products")

	fmt.Println(string(result))

	// myString := string(result)
	// fmt.Printf("%T", myString)
	// fmt.Println(myString)

}
