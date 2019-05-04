package main

import (
	"encoding/json"
	_ "log"
	_ "math/rand"
	"net/http"
	_ "strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     int     `json:"book_id"`
	Title  string  `json:"book_title"`
	Author *Author `json:"book_author"`
}

//Author Struct (Model)
type Author struct {
	ID       int    `json:"author_id"`
	FirtName string `json:"author_first_name"`
	LastName string `json:"author_last_name"`
}

var theBooks []Book

func main() {
	// Init the router
	r := mux.NewRouter()

	// Adding some data
	theBooks = append(theBooks, Book{ID: 1, Title: "Clean Code", Author: &Author{ID: 1, FirtName: "Robert Cecil", LastName: "Martin"}})
	theBooks = append(theBooks, Book{ID: 2, Title: "The Clean Coder", Author: &Author{ID: 1, FirtName: "Robert Cecil", LastName: "Martin"}})
	theBooks = append(theBooks, Book{ID: 3, Title: "Web Design with HTML, CSS, Javascript and Jquery Set", Author: &Author{ID: 2, FirtName: "Jon", LastName: "Duckett"}})
	theBooks = append(theBooks, Book{ID: 4, Title: "Data Science from Scratch - First Principles with Python", Author: &Author{ID: 1, FirtName: "Joel", LastName: "Grus"}})

	settingTheRouter(r)

	http.ListenAndServe(":8000", r)
}

func settingTheRouter(routeGuy *mux.Router) {

	// Route Handlers
	routeGuy.HandleFunc("/api/v1/books", getBooks).Methods("GET")

	routeGuy.HandleFunc("/api/v1/books/{id}", getBook).Methods("GET")

	routeGuy.HandleFunc("/api/v1/books", createBooks).Methods("POST")

	routeGuy.HandleFunc("/api/v1/books/{id}", updateBooks).Methods("PUT")

	routeGuy.HandleFunc("/api/v1/books/{id}", deleteBooks).Methods("DELETE")
}

func getBooks(writerGuy http.ResponseWriter, requestGuy *http.Request) {

	writerGuy.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writerGuy).Encode(theBooks)
}

func createBooks(writerGuy http.ResponseWriter, requestGuy *http.Request) {

}

func updateBooks(writerGuy http.ResponseWriter, requestGuy *http.Request) {

}

func deleteBooks(writerGuy http.ResponseWriter, requestGuy *http.Request) {

}

func getBook(writerGuy http.ResponseWriter, requestGuy *http.Request) {

}
