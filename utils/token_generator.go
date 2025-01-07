package utils

import (
	"github.com/dgrijalva/jwt-go"
	ws "github.com/yebology/giggle-backend/model/http"
)

func GenerateJWT(user ws.User) (string, error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
		"email": user.Email,
		"username": user.Username,
		"role": user.Role,
	})

	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}

	return signedToken, err

}