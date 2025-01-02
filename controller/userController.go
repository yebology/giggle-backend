package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/global"
	"github.com/yebology/giggle-backend/model"
	"github.com/yebology/giggle-backend/model/data"
	"github.com/yebology/giggle-backend/output"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var post model.Post
	err := c.BodyParser(&post)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	err = global.GetValidator().Struct(post)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	if post.PostType == "Hire" && post.RequiredTalent == 0 {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.HirePostError))
	}

	objectId, err := primitive.ObjectIDFromHex(post.PostCreatorId.Hex())
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}
	post.PostCreatorId = objectId

	collection := database.GetDatabase().Collection("post")
	_, err = collection.InsertOne(ctx, post)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully created a new post!",
		"data": fiber.Map{
			"post": post,
		},
	})

}

func UpdatePost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	var post model.Post
	err = c.BodyParser(&post)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	collection := database.GetDatabase().Collection("post")
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": post}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToUpdateData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully updated a post!",
		"data": fiber.Map{
			"post": post,
		},
	})

}

func DeletePost(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id) 
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	collection := database.GetDatabase().Collection("post")
	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToDeleteData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Succesfully deleted a post!",
		"data": "",
	})

}

func CreateGroup(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var group model.Group
	err := c.BodyParser(&group)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	_, err = primitive.ObjectIDFromHex(group.GroupOwnerId.Hex())
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	collection := database.GetDatabase().Collection("group")
	_, err = collection.InsertOne(ctx, group)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully created a new group!",
		"data": fiber.Map{
			"group": group,
		},
	})

}

func InviteMemberToGroup(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	var invitation data.Invitation
	err = c.BodyParser(&invitation)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	_, err = primitive.ObjectIDFromHex(invitation.MemberId.Hex())
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	var group model.Group
	filter := bson.M{"_id": objectId}

	collection := database.GetDatabase().Collection("group")
	err = collection.FindOne(ctx, filter).Decode(&group)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToDecodeData))
	}

	groupMemberIds := append(group.GroupMemberIds, invitation.MemberId)
	update := bson.M{
		"$set": bson.M{
			"_groupMemberIds": groupMemberIds,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToUpdateData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully invited new member!",
		"data": fiber.Map{
			"group": group,
		},
	})

}

func GetUserGroups(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("user_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	var groups []model.Group
	collection := database.GetDatabase().Collection("group")
	filter := bson.M{"groupOwnerId": objectId}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToRetrieveData))
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &groups)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToDecodeData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully fetched user groups!",
		"data": fiber.Map{
			"groups": groups,
		},
	})

}