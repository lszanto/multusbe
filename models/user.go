package models

import "github.com/lszanto/multusbe/multus"

// User ma man
type User struct {
	multus.Model
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"-" validate:"required"`
}
