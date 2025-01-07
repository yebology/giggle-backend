package helper

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func ParseToken(c *fiber.Ctx) (jwt.MapClaims, error) {

	token := c.Get("Authorization")
	if token == "" {
		token = c.Query("Authorization")
	}	

	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, GetSecretKey)
	if err != nil || !parsedToken.Valid {
		return nil, err
	}

	return claims, nil

}

func GetSecretKey(token *jwt.Token) (interface{}, error) {

	return []byte("secret-key"), nil

}