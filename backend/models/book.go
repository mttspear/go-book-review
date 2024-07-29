package models

import (
	"database/sql"
)

// Book represents the structure of a book in the system
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// CreateBook inserts a new book record into the database
func CreateBook(db *sql.DB, book *Book) error {
	query := `INSERT INTO books (title, author, book_description) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, book.Title, book.Author, book.Description)
	return err
}

// GetBook retrieves a book by ID
func GetBookByID(db *sql.DB, id int) (*Book, error) {
	query := `SELECT id, title, author, book_description FROM books WHERE id = $1`
	row := db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// GetBooks retrieves all books
func GetAllBooks(db *sql.DB) ([]Book, error) {
	query := `SELECT id, title, author, book_description FROM books`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// UpdateBook updates an existing book record
func UpdateBook(db *sql.DB, book *Book) error {
	query := `UPDATE books SET title = $1, author = $2, description = $3 WHERE id = $4`
	_, err := db.Exec(query, book.Title, book.Author, book.Description, book.ID)
	return err
}

// DeleteBook deletes a book record by ID
func DeleteBook(db *sql.DB, id int) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}
