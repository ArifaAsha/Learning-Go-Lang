package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	ID         int    `json:id`
	Name       string `json:name`
	Department string `json:department`
	DOB        string `json:dob`
}

var students []Student

func main() {
	router := mux.NewRouter()

	students = append(students, Student{ID: 1, Name: "A. Rahim", Department: "CSE", DOB: "09-09-1998"},
		Student{ID: 2, Name: "A. Karim", Department: "EEE", DOB: "01-09-1997"})

	router.HandleFunc("/students", getStudents).Methods("GET") //Method-> GET action
	router.HandleFunc("/student/{id}", getStudent).Methods("GET")
	router.HandleFunc("/students", addStudent).Methods("POST")
	router.HandleFunc("/students", updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", removeStudent).Methods("DELETE ")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	// log.Println("Gets all books")
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	// log.Println("Gets a book")
	params := mux.Vars(r) //returns a map

	log.Println(params)

	// log.Println(reflect.TypeOf(params["id"]))

	i, _ := strconv.Atoi(params["id"])

	for _, student := range students {
		if student.ID == i {
			json.NewEncoder(w).Encode(&student)
		}
	}
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("Add a student")
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("Updates a student")
}

func removeStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("Removes a student")
}
