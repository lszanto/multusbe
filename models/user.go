package models

import "github.com/lszanto/multusbe/multus"

// User ma man
type User struct {
	multus.Model
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"email" gorm:"unique"`
	Password string `json:"password" validate:"required"`
}
