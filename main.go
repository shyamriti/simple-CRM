package main

import (
	"simpleCRM/pkg/controller"
	"simpleCRM/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func setupRouters(app *fiber.App) {
	app.Get("/person/:id", controller.GetPerson)
	app.Get("/person", controller.GetPersons)
	app.Post("/person", controller.NewPerson)
	app.Delete("/person/:id", controller.DeletePerson)
}
func main() {
	app := fiber.New()
	database.DbConn()
	setupRouters(app)
	app.Listen(":8000")
}
