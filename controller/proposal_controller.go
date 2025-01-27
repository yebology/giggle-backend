package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/giggle-backend/constant"
	"github.com/yebology/giggle-backend/database"
	"github.com/yebology/giggle-backend/global"
	"github.com/yebology/giggle-backend/model/http"
	"github.com/yebology/giggle-backend/output"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProposal(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var proposal http.Proposal
	err := c.BodyParser(&proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	err = global.GetValidator().Struct(proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.ValidationError))
	}

	_, err = primitive.ObjectIDFromHex(proposal.PostId.Hex())
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	_, err = primitive.ObjectIDFromHex(proposal.CreatorId.Hex())
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	// _, err = primitive.ObjectIDFromHex(proposal.BuyerId.Hex())
	// if err != nil {
	// 	return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	// }

	collection := database.GetDatabase().Collection("proposal")
	_, err = collection.InsertOne(ctx, proposal)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully created a new proposal!",
		"data": "",
	})

}

func AcceptProposal(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	collection := database.GetDatabase().Collection("proposal")
	filter := bson.M{"_id": objectId}
	update := bson.M{"acceptByBuyer": true}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToUpdateData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully updated a proposal status!",
		"data": "",
	})

}

func GetProposals(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := c.Params("user_id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.InvalidIdError))
	}

	var proposals []http.Proposal

	collection := database.GetDatabase().Collection("proposal")
	filter := bson.M{
		"$or": bson.M{
			"creatorId": objectId,
			"buyerId": objectId,
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToRetrieveData))
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &proposals)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToDecodeData))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully fetched user's proposals!",
		"proposals": proposals,
	})

}