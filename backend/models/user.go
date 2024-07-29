package models

import (
	"database/sql"
	"errors"
	"time"
)

// User represents the user data structure
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, user *User) error {
	query := `
        INSERT INTO users (email, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	err := db.QueryRow(query, user.Email, user.Password, time.Now(), time.Now()).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE id = $1`

	row := db.QueryRow(query, id)

	var user User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by their email
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`

	row := db.QueryRow(query, email)

	var user User
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user's details
func UpdateUser(db *sql.DB, user *User) error {
	query := `
        UPDATE users
        SET email = $1, password = $2, updated_at = $3
        WHERE id = $4`

	res, err := db.Exec(query, user.Email, user.Password, time.Now(), user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// DeleteUser deletes a user by their ID
func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
