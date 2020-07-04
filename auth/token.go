package auth

import (
	"coffeebeans-people-backend/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (svc *Service) GenerateToken(user *models.User) (string, error) {

	claims := models.JWTTokenClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "coffeebeans-people",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(svc.SecretKey))

	return tokenString, err
}
