package models

import (
	"database/sql"
	"errors"
	"time"
)

// Review represents the review data structure
type Review struct {
	ID        int       `json:"id"`
	BookID    int       `json:"book_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateReview inserts a new review into the database
func CreateReview(db *sql.DB, review *Review) error {
	query := `
        INSERT INTO reviews (book_id, user_id, content, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`

	err := db.QueryRow(query, review.BookID, review.UserID, review.Content, time.Now(), time.Now()).Scan(&review.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetAllReviews retrieves all reviews from the database
func GetAllReviews(db *sql.DB) ([]Review, error) {
	query := `SELECT id, book_id, user_id, content, created_at, updated_at FROM reviews`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		if err := rows.Scan(&review.ID, &review.BookID, &review.UserID, &review.Content, &review.CreatedAt, &review.UpdatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// GetReviewByID retrieves a review by its ID
func GetReviewByID(db *sql.DB, id int) (*Review, error) {
	query := `SELECT id, book_id, user_id, content, created_at, updated_at FROM reviews WHERE id = $1`

	row := db.QueryRow(query, id)

	var review Review
	if err := row.Scan(&review.ID, &review.BookID, &review.UserID, &review.Content, &review.CreatedAt, &review.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("review not found")
		}
		return nil, err
	}
	return &review, nil
}

// UpdateReview updates an existing review
func UpdateReview(db *sql.DB, review *Review) error {
	query := `
        UPDATE reviews
        SET book_id = $1, user_id = $2, content = $3, updated_at = $4
        WHERE id = $5`

	res, err := db.Exec(query, review.BookID, review.UserID, review.Content, time.Now(), review.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("review not found")
	}
	return nil
}

// DeleteReview deletes a review by its ID
func DeleteReview(db *sql.DB, id int) error {
	query := `DELETE FROM reviews WHERE id = $1`

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("review not found")
	}
	return nil
}
