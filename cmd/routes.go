package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wbrijesh/console-go/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/posts", handlers.ListPosts)
	app.Post("/posts", handlers.CreatePost)
	app.Delete("/posts/:id", handlers.DeletePost)
	app.Put("/posts/:id", handlers.EditPost)
}
