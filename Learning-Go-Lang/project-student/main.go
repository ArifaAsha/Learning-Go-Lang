package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	//connect to DB istance for elephant sql
	"github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

type Student struct {
	ID         int    `json:id`
	Name       string `json:name`
	Department string `json:department`
	DOB        string `json:dob`
}

var students []Student
var db *sql.DB

func init() {
	gotenv.Load() //loads all the env variable inside the .env file
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)

	logFatal(err)

	err = db.Ping()
	logFatal(err) //if ping() gives error

	// log.Println(pgURL)

	/*
		pgURL gives:
		-`dbname`
		-`host`
		-`password`
		-`port`
		-`user`
		pass into sql.open()
	*/

	router := mux.NewRouter()

	router.HandleFunc("/students", getStudents).Methods("GET") //Method-> GET action
	router.HandleFunc("/students/{id}", getStudent).Methods("GET")
	router.HandleFunc("/students", addStudent).Methods("POST")
	router.HandleFunc("/students", updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", removeStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	var student Student
	students = []Student{}
	rows, err := db.Query("select * from student order by id")
	logFatal(err)

	defer rows.Close()

	for rows.Next() { //map
		err := rows.Scan(&student.ID, &student.Name, &student.Department, &student.DOB)
		logFatal(err) //if error

		students = append(students, student)
	}

	json.NewEncoder(w).Encode(students)

}

func getStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	params := mux.Vars(r)

	rows := db.QueryRow("select * from student where id=$1", params["id"]) //$1 => placeholder id inside params
	err := rows.Scan(&student.ID, &student.Name, &student.Department, &student.DOB)
	logFatal(err)

	json.NewEncoder(w).Encode(student)
}

func addStudent(w http.ResponseWriter, r *http.Request) {

}

func updateStudent(w http.ResponseWriter, r *http.Request) {

}

func removeStudent(w http.ResponseWriter, r *http.Request) {

}
