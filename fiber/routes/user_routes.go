package routes

import (
	"github.com/api-rest-fiber/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoutes(app *fiber.App, collection *mongo.Collection) {

	app.Get("/users", func(ctx *fiber.Ctx) error {
		results, err := services.GetAllUsers(collection)
		if err != nil {
			return ctx.JSON(&fiber.Map{
				"error": err,
			})
		}
		return ctx.JSON(&fiber.Map{
			"data": results,
		})
	})

	app.Post("/users", func(ctx *fiber.Ctx) error {
		result, err := services.CreateUser(ctx, collection)

		if err != nil {
			return ctx.JSON(&fiber.Map{
				"error": err,
			})
		}
		return ctx.JSON(&fiber.Map{
			"data": result,
		})
	})
}
