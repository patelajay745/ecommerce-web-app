package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("../templates", ".html"),
	})

	app.Static("/", "../static")
	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("login/signup", fiber.Map{})
	})

	app.Post("/signup", func(c *fiber.Ctx) error {
		fmt.Println(c.FormValue("fname"))
		fmt.Println(c.FormValue("gender"))
		return nil
	})

	app.Listen(":3000")
}
