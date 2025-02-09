package routes

import (
	"html/template"
	"net/http"

	"inventory-book/controllers"
	"inventory-book/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// HTML Routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/books.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	}).Methods("GET")

	r.HandleFunc("/books", controllers.GetBooksPage).Methods("GET")

	r.HandleFunc("/login", controllers.LoginPage).Methods("GET")

	r.HandleFunc("/register", controllers.RegisterPage).Methods("GET")

	// Public API Routes
	publicAPI := r.PathPrefix("/api").Subrouter()
	publicAPI.HandleFunc("/register", controllers.Register).Methods("POST")
	publicAPI.HandleFunc("/login", controllers.Login).Methods("POST")

	// Protected API Routes
	protectedAPI := r.PathPrefix("/api").Subrouter()
	protectedAPI.Use(middlewares.AuthMiddleware)
	protectedAPI.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	protectedAPI.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	protectedAPI.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	protectedAPI.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	return r
}
