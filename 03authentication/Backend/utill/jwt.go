package utill

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte("MY_SECRET_KEY")

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET)
}
