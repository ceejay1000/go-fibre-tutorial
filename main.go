package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Static("/views", "/public/views")
	app.Static("/", "/public")

	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "Some custom value")
		log.Println("Sets a custom header here")
		return c.Next()
	}, func(c *fiber.Ctx) error {
		log.Println("Does nothing here")
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		log.Println("Logging for all requests => " + c.Method())
		return c.Next()
	})

	app.Use("/api", func(c *fiber.Ctx) error {
		log.Println("Logging for all /api requests => " + c.Method())
		return c.Next()
	})

	app.Use([]string{"/api", "/public"}, func(c *fiber.Ctx) error {
		log.Println("Logging for all /api and /public requests => " + c.Method())
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// c.Response().Header.Set(fmt.Sprint("%i", fiber.StatusAccepted), "Status ok")

		return c.SendString("Hello")
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("name"))
	})

	app.Add(fiber.MethodGet, "/add", func(c *fiber.Ctx) error {
		return c.SendString("Add route")
	})

	app.All("/all-routes", func(c *fiber.Ctx) error {
		return c.SendString("All routes")
	})

	app.Get("/:age<int>", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/:active<bool>", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/:username<minLen(4)>", func(c *fiber.Ctx) error {
		return nil
	})

	app.Post("/submission-date:date<regex(\\d{4}-\\d{2}-\\d{2})}>", func(c *fiber.Ctx) error {
		return nil
	})

	log.Fatalln(app.Listen(":8082"))
}
