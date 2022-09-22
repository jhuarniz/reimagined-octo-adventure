package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jhuarniz/reimagined-octo-adventure/api/models"
)

func UserController(router fiber.Router) {
	userGroup := router.Group("/users")
	userGroup.Get("", user)
	userGroup.Post("", createUser)
}

func user(c *fiber.Ctx) error {
	user := models.User{
		Firstname: "Joe",
		Lastname:  "Doe",
		Age:       30,
	}
	//return c.Status(fiber.StatusOK).JSON(user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func createUser(c *fiber.Ctx) error {
	fmt.Println("Users - Post")
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	users := []models.User{user}
	return c.Status(fiber.StatusOK).JSON(users)
}
