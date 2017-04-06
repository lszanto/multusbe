package models

import "github.com/lszanto/multusbe/multus"

// Doc ma man
type Doc struct {
	multus.Model
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
