package lib

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(jwtToken string, jwtKey []byte) (*claims, bool) {
	claims := &claims{}
	_, err := jwt.ParseWithClaims(jwtToken, claims, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		return nil, false
	}

	return claims, true
}

func NewSignedToken(id string, key []byte, duration time.Duration) (string, error) {
	signed, err := NewToken(id, duration).SignedString(key)
	if err != nil {
		return "", err
	}
	return signed, nil
}

type claims struct {
	ID string
	jwt.StandardClaims
}

func NewToken(id string, duration time.Duration) *jwt.Token {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&claims{
			ID: id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(duration).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		})

	return token
}
