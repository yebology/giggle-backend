package helper

import (
	"context"

	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func CheckUser(ctx context.Context, filter bson.M) (model.User, error) {

	var user model.User
	collection := database.GetDatabase().Collection("user")

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return model.User{},err
	}

	return user, nil

}

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil

}

func CheckPassword(hashedPassword string, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))	

}
