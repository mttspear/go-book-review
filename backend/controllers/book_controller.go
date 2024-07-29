package controllers

import (
	"encoding/json"
	"go-book-review/models"
	"go-book-review/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateBook handles the creation of a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Save the book to the database
	if err := models.CreateBook(utils.GetDB(), &book); err != nil {
		http.Error(w, "Error creating book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetAllBooks retrieves all books from the database
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks(utils.GetDB())
	if err != nil {
		http.Error(w, "Error retrieving books", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

// GetBookByID retrieves a book by its ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	bookIDStr := mux.Vars(r)["id"]
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := models.GetBookByID(utils.GetDB(), bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// UpdateBook updates an existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookIDStr := mux.Vars(r)["id"]
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	book.ID = bookID

	if err := models.UpdateBook(utils.GetDB(), &book); err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// DeleteBook deletes a book by its ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookIDStr := mux.Vars(r)["id"]
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if err := models.DeleteBook(utils.GetDB(), bookID); err != nil {
		http.Error(w, "Error deleting book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
