package controllers

import (
	"encoding/json"
	"go-book-review/models"
	"go-book-review/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateReview handles the creation of a new review
func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Extract user ID from context
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	review.UserID = userID

	// Save the review to the database
	if err := models.CreateReview(utils.GetDB(), &review); err != nil {
		http.Error(w, "Error creating review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

// GetReviews retrieves all reviews
func GetReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := models.GetAllReviews(utils.GetDB())
	if err != nil {
		http.Error(w, "Error fetching reviews", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reviews)
}

// GetReview retrieves a single review by ID
func GetReview(w http.ResponseWriter, r *http.Request) {
	reviewIDStr := mux.Vars(r)["id"]
	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	review, err := models.GetReviewByID(utils.GetDB(), reviewID)
	if err != nil {
		http.Error(w, "Review not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(review)
}

// UpdateReview updates an existing review by ID
func UpdateReview(w http.ResponseWriter, r *http.Request) {
	reviewIDStr := mux.Vars(r)["id"]
	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	review.ID = reviewID

	// Update the review in the database
	if err := models.UpdateReview(utils.GetDB(), &review); err != nil {
		http.Error(w, "Error updating review", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(review)
}

// DeleteReview deletes a review by ID
func DeleteReview(w http.ResponseWriter, r *http.Request) {
	reviewIDStr := mux.Vars(r)["id"]
	reviewID, err := strconv.Atoi(reviewIDStr)
	if err != nil {
		http.Error(w, "Invalid review ID", http.StatusBadRequest)
		return
	}

	// Delete the review from the database
	if err := models.DeleteReview(utils.GetDB(), reviewID); err != nil {
		http.Error(w, "Error deleting review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
