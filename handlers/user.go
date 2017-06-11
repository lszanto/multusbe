package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/jinzhu/gorm"
	"github.com/lszanto/multusbe/models"
	"github.com/lszanto/multusbe/multus"
	"github.com/lszanto/multusbe/response"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler handles user activity
type UserHandler struct {
	DB     *gorm.DB
	Config multus.Config
}

// NewUserHandler creates a new user handler
func NewUserHandler(db *gorm.DB, config multus.Config) UserHandler {
	return UserHandler{DB: db, Config: config}
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
