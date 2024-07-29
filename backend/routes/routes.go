package routes

import (
	"go-book-review/controllers"
	"go-book-review/utils"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Authentication routes
	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Book routes
	r.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/books", utils.AuthMiddleware(controllers.CreateBook)).Methods("POST")
	r.HandleFunc("/books/{id}", utils.AuthMiddleware(controllers.UpdateBook)).Methods("PUT")
	r.HandleFunc("/books/{id}", utils.AuthMiddleware(controllers.DeleteBook)).Methods("DELETE")

	// Review routes
	r.HandleFunc("/books/{id}/reviews", controllers.GetReviews).Methods("GET")
	r.HandleFunc("/books/{id}/reviews", utils.AuthMiddleware(controllers.CreateReview)).Methods("POST")
	r.HandleFunc("/books/{id}/reviews/{reviewId}", utils.AuthMiddleware(controllers.UpdateReview)).Methods("PUT")
	r.HandleFunc("/books/{id}/reviews/{reviewId}", utils.AuthMiddleware(controllers.DeleteReview)).Methods("DELETE")

	// User routes
	r.HandleFunc("/users/{id}", utils.AuthMiddleware(controllers.GetUser)).Methods("GET")
	// r.HandleFunc("/users/{id}", utils.AuthMiddleware(controllers.UpdateUser)).Methods("PUT")
	// r.HandleFunc("/users/{id}", utils.AuthMiddleware(controllers.DeleteUser)).Methods("DELETE")

	return r
}
