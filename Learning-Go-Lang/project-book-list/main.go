package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang Pointers", Author: "Mr. Go", Year: "2010"},
		Book{ID: 2, Title: "Golang coroutines", Author: "Mr. Goroutine", Year: "2020"})

	router.HandleFunc("/books", getBooks).Methods("GET") //Method-> GET action
	router.HandleFunc("/books/{ID}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{ID}", removeBook).Methods("DELETE ")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// log.Println("Gets all books")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	// log.Println("Gets a book")

}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add one book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updates a book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Removes a book")
}
