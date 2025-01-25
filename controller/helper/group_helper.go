package helper

import (
	"context"
	"time"

	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/model/http"
	"go.mongodb.org/mongo-driver/bson"
)
// GetGroupByFilter retrieves all groups from the database that match a specific filter
func GetGroupByFilter(filter bson.M) ([]http.Group, error) {

	// Create a context with timeout to prevent blocking the request indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Initialize a slice to hold the groups that will be retrieved
	var groups []http.Group
	collection := database.GetDatabase().Collection("group")

	// Execute the find query with the provided filter
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		// If there is an error with the query, return an empty slice and the error
		return []http.Group{}, err
	}
	// Ensure that the cursor is closed once the operation is done
	defer cursor.Close(ctx)

	// Decode the results of the query into the groups slice
	err = cursor.All(ctx, &groups)
	if err != nil {
		// If there is an error while decoding the results, return an empty slice and the error
		return []http.Group{}, err
	}

	// Return the retrieved groups and nil for error if successful
	return groups, nil
	
}
