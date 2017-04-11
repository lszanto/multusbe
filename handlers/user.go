package handlers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/lszanto/multusbe/response"
)

// UserHandler handles user activity
type UserHandler struct {
	DB *gorm.DB
}

// NewUserHandler creates a new user handler
func NewUserHandler(db *gorm.DB) UserHandler {
	return UserHandler{DB: db}
}

// Login allows the user to login
func (handler UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, "hello", http.StatusOK)
}
