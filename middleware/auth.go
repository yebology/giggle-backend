package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/output"
)

func UserMiddleware(c *fiber.Ctx) error {

	claims, err := ParseToken(c)
	if err != nil {
		return output.GetError(c, fiber.StatusUnauthorized, "Invalid token!")
	}

	var expectedRole = "user"
	role, ok := claims["role"].(string)
	if role != expectedRole || !ok {
		return output.GetError(c, fiber.StatusForbidden, "Permission denied! Must register or login first!")
	}

	return c.Next()

}

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