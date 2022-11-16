package main

import (
	"fmt"
	"strings"
)

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func main() {
	// Example string to parse.
	test := "{\"product":{"id":8008759542065,"title":"Antique Drawers","body_html":"\u003cp\u003eAntique wooden chest of drawers\u003c\/p\u003e","vendor":"Company 123","product_type":"Indoor","created_at":"2022-11-16T09:53:10+06:00","handle":"antique-drawers","updated_at":"2022-11-16T09:53:21+06:00","published_at":"2022-11-16T09:53:09+06:00","template_suffix":null,"status":"active","published_scope":"global","tags":"Antique, Bedroom","admin_graphql_api_id":"gid:\/\/shopify\/Product\/8008759542065","variants":[{"id":44053623996721,"product_id":8008759542065,"title":"Default Title","price":"250.00","sku":null,"position":1,"inventory_policy":"deny","compare_at_price":"300.00","fulfillment_service":"manual","inventory_management":null,"option1":"Default Title","option2":null,"option3":null,"created_at":"2022-11-16T09:53:10+06:00","updated_at":"2022-11-16T09:53:10+06:00","taxable":true,"barcode":null,"grams":0,"image_id":null,"weight":0.0,"weight_unit":"kg","inventory_item_id":46101467070769,"inventory_quantity":2,"old_inventory_quantity":2,"requires_shipping":true,"admin_graphql_api_id":"gid:\/\/shopify\/ProductVariant\/44053623996721"}],"options":[{"id":10176645366065,"product_id":8008759542065,"name":"Title","position":1,"values":["Default Title"]}],"images":[{"id":39852047171889,"product_id":8008759542065,"position":1,"created_at":"2022-11-16T09:53:10+06:00","updated_at":"2022-11-16T09:53:10+06:00","alt":null,"width":925,"height":617,"src":"https:\/\/cdn.shopify.com\/s\/files\/1\/0679\/7639\/7105\/products\/babys-room_925x_ffcafdc2-b8f3-4ccc-b3be-159c634314ac.jpg?v=1668570790","variant_ids":[],"admin_graphql_api_id":"gid:\/\/shopify\/ProductImage\/39852047171889"}],"image":{"id":39852047171889,"product_id":8008759542065,"position":1,"created_at":"2022-11-16T09:53:10+06:00","updated_at":"2022-11-16T09:53:10+06:00","alt":null,"width":925,"height":617,"src":"https:\/\/cdn.shopify.com\/s\/files\/1\/0679\/7639\/7105\/products\/babys-room_925x_ffcafdc2-b8f3-4ccc-b3be-159c634314ac.jpg?v=1668570790","variant_ids":[],"admin_graphql_api_id":"gid:\/\/shopify\/ProductImage\/39852047171889"}}}"

		// Test between func.
		fmt.Println(between(test, "DEFINE:", "="))
	fmt.Println(between(test, ":", "="))

	// Test before func.
	fmt.Println(before(test, ":"))
	fmt.Println(before(test, "="))

	// Test after func.
	fmt.Println(after(test, ":"))
	fmt.Println(after(test, "DEFINE:"))
	fmt.Println(after(test, "="))
}
