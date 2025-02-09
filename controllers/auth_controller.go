package controllers

import (
	"encoding/json"
	"html/template"
	"inventory-book/config"
	"inventory-book/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoadTemplates(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	LoadTemplates(w, "register.html", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// Check request method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Request Body
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Save user to database
	config.DB.Create(&user)

	// Response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	LoadTemplates(w, "login.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Check request method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Request Body
	var input models.User
	var user models.User
	json.NewDecoder(r.Body).Decode(&input)

	// Check user
	config.DB.Where("username = ?", input.Username).First(&user)
	if user.ID == 0 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Verify password
	if passwordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); passwordErr != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Response
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
