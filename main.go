package main

import (
	"encoding/json"
	"log"
	"strconv"

	//"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jashfeer/RESTfulAPI/initialization"
	//	"gorm.io/gorm"
)

type Book struct {
	ID     int    `  json:"id" `
	Title  string `json:"title"`
	Price  string `json:"price"`
	Author string `json:"author"`
}

var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := initialization.Db.Find(&books)
	if result.Error != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(books)

}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	params := mux.Vars(r) // Gets params

	result := initialization.Db.Where("id = ?", params["id"]).Find(&book)
	if result.Error != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	json.NewEncoder(w).Encode(&book)
}

// Add new book

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	initialization.Db.Select("title", "price", "author").Create(&book)

	json.NewEncoder(w).Encode(book)

}

//Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	initialization.Db.Model(&book).Where("id = ?", params["id"]).Updates(Book{Title:book.Title, Price: book.Price, Author:book.Author})
	
	book.ID,_=  strconv.Atoi(params["id"])

	json.NewEncoder(w).Encode(book)
}

//Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	initialization.Db.Delete(&Book{},params["id"])
	 initialization.Db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()
	initialization.Init()

	// Route handles & endpoints

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	 r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	 r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	 

	log.Println("Runing...in port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
