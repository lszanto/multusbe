package handlers

import (
	"encoding/json"
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jinzhu/gorm"
	"github.com/lszanto/links/models"
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
		response.JSON(w, response.Result{Result: "Error creating user"}, http.StatusNotAcceptable)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	handler.DB.Create(&user)
	response.JSON(w, response.Result{Result: "User created"}, http.StatusCreated)
}

// Login allows the user to login
func (handler UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, "hello", http.StatusOK)
}
