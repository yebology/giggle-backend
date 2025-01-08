package helper

import (
	"context"
	"time"

	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model/http"
	"go.mongodb.org/mongo-driver/bson"
)

func GetGroupByFilter(filter bson.M) ([]http.Group, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var groups []http.Group
	collection := database.GetDatabase().Collection("group")

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []http.Group{}, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &groups)
	if err != nil {
		return []http.Group{}, err
	}

	return groups, nil

}