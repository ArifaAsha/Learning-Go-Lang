package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// type Podcast struct {
// 	ID     primitive.ObjectID `bson:"_id,omitempty"`
// 	Title  string             `bson:"title,omitempty"`
// 	Author string             `bson:"author,omitempty"`
// 	Tags   []string           `bson:"tags,omitempty"`
// }

type Item struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

func webScrapper() {
	items := []Item
	c := colly.NewCollector(colly.AllowedDomains("ratiocoffee.com"))
	c.OnHTML("div.grid grid--uniform", func(element *colly.HMLElement) {
		products := element.DOM
		product := Item{
			Name:  products.Find("div.grid-product__title").Text(),
			Price: products.Find("div.grid-product__price").Text(),
		}
		items = append(products, product)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})
}

// func main() {

	

// }

// func main() {
// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://asha:lDFQuDlbGR8RMFwr@cluster0.ebrqthb.mongodb.net/shopify?retryWrites=true&w=majority"))
// if err != nil {
// 	log.Fatal(err)
// }
// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// err = client.Connect(ctx)
// if err != nil {
// 	log.Fatal(err)
// }
// defer client.Disconnect(ctx)

// err = client.Ping(ctx, readpref.Primary())
// if err != nil {
// 	log.Fatal(err)
// }

// databases, err := client.ListDatabaseNames(ctx, bson.M{})
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(databases)

// podcast := Podcast{
// 	Title:  "The Polyglot Developer",
// 	Author: "Nic Raboy",
// 	Tags:   []string{"development", "programming", "coding"},
// }

// database := client.Database("shopify")
// podcastsCollection := database.Collection("shopifyData")

// insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
// if err != nil {
// 	panic(err)
// }
// fmt.Println(insertResult.InsertedID)

// }
