package helper

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func ParseToken(c *fiber.Ctx) (jwt.MapClaims, error) {

	token := c.Get("Authorization")

	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, GetSecretKey)
	fmt.Println(claims)
	fmt.Println(parsedToken)
	if err != nil || !parsedToken.Valid {
		fmt.Println(err)
		fmt.Println(parsedToken.Valid)
		return nil, err
	}

	return claims, nil

}

func GetSecretKey(token *jwt.Token) (interface{}, error) {

	return []byte("secret-key"), nil

}