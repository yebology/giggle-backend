package helper

import (
	"context"

	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(ctx context.Context, filter bson.M) (model.User, error) {

	var user model.User
	collection := database.GetDatabase().Collection("user")
	
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return model.User{}, err
	}
	defer cursor.Close(ctx)

	return user, nil

}