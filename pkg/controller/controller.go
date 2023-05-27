package controller

import (
	"fmt"
	"simpleCRM/pkg/database"
	"simpleCRM/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func GetPerson(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Db
	var person models.Person
	db.Find(&person, id)
	c.JSON(person)
	return nil
}

func GetPersons(c *fiber.Ctx) error {
	var persons []models.Person
	db := database.Db
	db.Find(&persons)
	c.JSON(persons)
	return nil
}

func NewPerson(c *fiber.Ctx) error {
	db := database.Db
	person := new(models.Person)
	if err := c.BodyParser(person); err != nil {

		fmt.Printf("err: %v\n", err)
	}
	db.Create(&person)
	c.JSON(person)
	return nil
}

func DeletePerson(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Db
	var person models.Person
	db.Delete(&person, id)
	c.Send([]byte("data deleted"))
	return nil
}
