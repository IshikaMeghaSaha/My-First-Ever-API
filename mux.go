package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

var Books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

func getBookbyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	s := p["id"]
	var book Book
	for _, book = range Books {
		if book.ID == s {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	Books = append(Books, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	s := p["id"]
	var index int
	var book Book
	for index, book = range Books {
		if book.ID == s {
			Books = append(Books[:index], Books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			Books = append(Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)
	s := p["id"]
	var index int
	var book Book
	for index, book = range Books {
		if book.ID == s {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Books)
}

func main() {
	router := mux.NewRouter()
	Books = append(Books, Book{ID: "1", Title: "Pride and Prejudice", Author: &Author{Name: "Jane Austen", Genre: "Romance"}})
	Books = append(Books, Book{ID: "2", Title: "Macbeth", Author: &Author{Name: "Shakespeare", Genre: "Tragedy"}})
	Books = append(Books, Book{ID: "3", Title: "Inferno", Author: &Author{Name: "Dan Brown", Genre: "Thriller"}})
	router.HandleFunc("/mux/Books", getBooks).Methods("GET")
	router.HandleFunc("/mux/Books/{id}", getBookbyID).Methods("GET")
	router.HandleFunc("/mux/Books", createBook).Methods("POST")
	router.HandleFunc("/mux/Books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/mux/Books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
