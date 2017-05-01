package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jinzhu/gorm"
	"github.com/lszanto/multusbe/models"
	"github.com/lszanto/multusbe/response"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler handles user activity
type UserHandler struct {
	DB *gorm.DB
}

// NewUserHandler creates a new user handler
func NewUserHandler(db *gorm.DB) UserHandler {
	return UserHandler{DB: db}
}

// Create makes a new user
func (handler UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		response.JSON(w, response.Result{Error: "Error creating user"}, http.StatusNotAcceptable)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	if err := handler.DB.Create(&user).Error; err != nil {
		errorMessage := "Error creating user"

		// if we've got a duplicate entry from the email then update message accordingly
		if strings.Contains(err.Error(), "Duplicate entry") {
			errorMessage = "A user with this email already exists"
		}

		response.JSON(w, response.Result{Error: errorMessage}, http.StatusNotAcceptable)
		return
	}
	response.JSON(w, response.Result{Result: "User created"}, http.StatusCreated)
}

// Login allows the user to login
func (handler UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var attempt models.User
	json.NewDecoder(r.Body).Decode(&attempt)

	var user models.User

	if err := handler.DB.Where("username = ?", attempt.Username).First(&user).Error; err != nil {
		response.JSON(w, response.Result{Error: "Could not login"}, http.StatusNotFound)
		return
	}

	// check that hashes match
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(attempt.Password)); err != nil {
		response.JSON(w, response.Result{Error: "Coult not login"}, http.StatusNotFound)
		return
	}

	response.JSON(w, response.Result{Result: "Logged in"}, http.StatusCreated)
}
