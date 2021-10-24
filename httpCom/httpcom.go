// https://github.com/bradtraversy/go_restapi

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/junwookheo/goexamples/httpCom/dbModel"
)

// Init books var as a slice Book struct
//var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbModel.GetBooks())
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	// for _, item := range books {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Book{})
	id, _ := strconv.Atoi(params["id"])
	json.NewEncoder(w).Encode(dbModel.GetBook(id))
}

// Add new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book dbModel.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	// books = append(books, book)
	b := dbModel.AddBook(book)
	json.NewEncoder(w).Encode(b)
}

// Update book
// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// }

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// for index, item := range books {
	// 	if item.ID == params["id"] {
	// 		books = append(books[:index], books[index+1:]...)
	// 		break
	// 	}
	// }
	// json.NewEncoder(w).Encode(books)

	id, _ := strconv.Atoi(params["id"])
	dbModel.RemoveBook(id)
	json.NewEncoder(w).Encode(dbModel.GetBooks())
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	// books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	dbModel.AddBook(dbModel.Book{ID: 1, Isbn: "438227", Title: "Book One", Author: &dbModel.Author{Firstname: "John", Lastname: "Doe"}})
	dbModel.AddBook(dbModel.Book{ID: 2, Isbn: "454555", Title: "Book Two", Author: &dbModel.Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	// r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
