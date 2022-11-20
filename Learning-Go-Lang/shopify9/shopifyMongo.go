package main

import (
	"context"
	"fmt"
	"log"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/rapito/go-shopify/shopify"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	store  = "dummy-store-abcd"
	apiKey = "7c92182fcb56bd9232b3f3771f9b8ffd"
	pass   = "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"
)

// Create a new shopify object with your store
// domain, api key and password
var shop = shopify.New(store, apiKey, pass)

// type Podcast struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty"`
// 	Title     string             `bson:"title,omitempty"`
// 	Body_html string             `bson:"body_html,omitempty"`
// 	Vendor    string             `bson:"vendor,omitempty"`
// 	Product_type product_type
// 	created_at
// 	handle
// 	updated_at
// 	published_at
// 	template_suffix
// 	status
// 	published_scope
// 	tags
// 	admin_graphql_api_id
// 	variants
// 	product_id

// }

// type Item struct {
// 	Name  string `json:"name"`
// 	Price string `json:"price"`
// }

// type Podcast struct {
// 	ID     primitive.ObjectID `bson:"_id,omitempty"`
// 	Title  string             `bson:"title,omitempty"`
// 	Author string             `bson:"author,omitempty"`
// 	Tags   []string           `bson:"tags,omitempty"`
// }

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

	products := jsonData.Get("products")

	// for k, v := range products.MustMap() {
	// 	fmt.Println("**************************")
	// 	fmt.Println(k, v)
	// 	fmt.Println("**************************")
	// }

	product := products.GetIndex(0)

	title, _ := product.Get("id").String()

	// fmt.Println("Full result: ")
	// fmt.Println("====================")

	// fmt.Println(string(result))

	// fmt.Println("First Product title: ")

	// fmt.Println("++++++++++++++++++++++++")
	fmt.Println(title)
	// mongodb+srv://asha:<password>@cluster0.ebrqthb.mongodb.net/?retryWrites=true&w=majority

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://asha:lDFQuDlbGR8RMFwr@cluster0.ebrqthb.mongodb.net/shopify?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	podcast := product
	database := client.Database("shopify")
	podcastsCollection := database.Collection("shopifyData")

	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
}
