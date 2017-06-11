package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/lszanto/multusbe/models"
	"github.com/lszanto/multusbe/multus"
	"github.com/lszanto/multusbe/response"
	"golang.org/x/crypto/bcrypt"
)

// AuthHandler handles user activity
type AuthHandler struct {
	DB     *gorm.DB
	Config multus.Config
}

// NewAuthHandler creates a new user handler
func NewAuthHandler(db *gorm.DB, config multus.Config) AuthHandler {
	return AuthHandler{DB: db, Config: config}
}

// Login allows the user to login
func (handler AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var attempt models.User
	json.NewDecoder(r.Body).Decode(&attempt)

	var user models.User

	if err := handler.DB.Where("username = ?", attempt.Username).First(&user).Error; err != nil {
		response.JSON(w, response.Result{Error: "Could not login"}, http.StatusNotFound)
		return
	}

	// check that hashes match
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(attempt.Password)); err != nil {
		response.JSON(w, response.Result{Error: "Could not login"}, http.StatusNotFound)
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, multus.CustomClaims{
		user.Email,
		jwt.StandardClaims{
			Id:        fmt.Sprint(user.ID),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(handler.Config.SecretKey)

	if err != nil {
		response.JSON(w, response.Result{Error: "Coult not login"}, http.StatusBadRequest)
		return
	}

	response.JSON(w, response.LoginResult{Token: tokenString}, http.StatusOK)
}
