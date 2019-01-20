package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Init Router
	r := mux.NewRouter()

	//Mock Data - @Todo - implement db
	//books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book one", Author: &Author{Firstname: "Ben", Lastname: "Dover"}})
	//books = append(books, Book{ID: "2", Isbn: "528721", Title: "Book two", Author: &Author{Firstname: "Aneta", Lastname: "Gofradump"}})
	connStr := "user=dev dbname=postgres-dev sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
