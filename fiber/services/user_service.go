package services

import (
	"context"
	"github.com/api-rest-fiber/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx *fiber.Ctx, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	var user models.User

	err := ctx.BodyParser(&user)
	if err != nil {
		return nil, err
	}

	result, err := collection.InsertOne(context.TODO(), bson.D{
		{
			Key:   "name",
			Value: user.Name,
		},
		{
			Key:   "email",
			Value: user.Email,
		},
	})

	return result, err
}

func GetAllUsers(collection *mongo.Collection) ([]models.User, error) {
	var users []models.User

	results, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		panic(err)
	}

	for results.Next(context.TODO()) {
		var user models.User
		err := results.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}
