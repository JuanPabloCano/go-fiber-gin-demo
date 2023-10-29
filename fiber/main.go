package main

import (
	"fmt"
	"github.com/api-rest-fiber/database"
	"github.com/api-rest-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	client := database.Connection()

	collection := client.Database("go-fiber-mongo").Collection("users")

	app.Use(cors.New())

	routes.UserRoutes(app, collection)

	err := app.Listen(":" + port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server listening on port 8080")
}
