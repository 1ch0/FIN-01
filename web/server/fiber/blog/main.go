package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Create a new endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"data": "Hello World!"})
	})
	setupRoutes(app)
	// Start server on port 3000
	app.Listen(":3000")
}

type Blog struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type blogPosts []Blog

var blogs = blogPosts{
	{Id: 1, Title: "Hello", Body: "Hello World!"},
	{Id: 2, Title: "Fiber", Body: "Fiber is fast!"},
	{Id: 3, Title: "Microservice", Body: "Microservice is awesome!"},
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/post", GetPosts)
	app.Get("/api/v1/post/:id", GetPost)
}

func GetPosts(c *fiber.Ctx) error {
	return c.JSON(blogs)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(fiber.Map{"error": "Invalid ID"})
	}
	for _, s := range blogs {
		if s.Id == i {
			return c.JSON(s)
		}
	}
	return c.JSON(fiber.Map{"error": "Post not found"})
}
