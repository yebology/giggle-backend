package helper

import (
	"context"

	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(ctx context.Context, filter bson.M) (model.User, error) {

	var user model.User
	collection := database.GetDatabase().Collection("user")
	
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return model.User{}, err
	}
	defer cursor.Close(ctx)

	success := cursor.Next(ctx)
	if !success {
		return model.User{}, err
	}

	err = cursor.Decode(&user)
	if err != nil {
		return model.User{}, err
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
