package controllers

import (
	"encoding/json"
	"html/template"
	"inventory-book/config"
	"inventory-book/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooksPage(w http.ResponseWriter, r *http.Request) {
	// Get all books
	var books []models.Book
	config.DB.Find(&books)

	t, err := template.ParseFiles("templates/books.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	t.Execute(w, books)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Get all books
	var books []models.Book
	config.DB.Find(&books)

	// Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Request body
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	// Save book to database
	config.DB.Create(&book)

	// Response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Request params
	id := mux.Vars(r)["id"]
	var book models.Book

	// Check if book exists in database
	if config.DB.First(&book, id).Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Request body
	json.NewDecoder(r.Body).Decode(&book)

	// Update book in database
	config.DB.Save(&book)

	// Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Request params
	id := mux.Vars(r)["id"]
	var book models.Book

	// Check if book exists in database
	if config.DB.First(&book, id).Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Delete book from database
	config.DB.Delete(&book)
	w.WriteHeader(http.StatusNoContent)
}
