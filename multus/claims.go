package multus

import (
	"github.com/dgrijalva/jwt-go"
)

// CustomClaims struct
type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
