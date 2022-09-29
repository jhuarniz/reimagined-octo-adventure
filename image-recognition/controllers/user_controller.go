package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jhuarniz/reimagined-octo-adventure/api/dtos"
	"github.com/thoas/go-funk"
)

func UserController(router fiber.Router) {
	userGroup := router.Group("/users")
	//userGroup.Get("", user)
	//userGroup.Post("", createUser)

	userGroup.Get("/", getUsers)
	userGroup.Get("/:id", getUser)
	userGroup.Post("/", createUser)
	//userGroup.Path("/:id", paritalUpdateUser)
	userGroup.Put("/:id", updateUser)
	userGroup.Delete("/:id", deleteUser)
}

var users = []*dtos.User{
	{
		Id:        "A001",
		Firstname: "Elmer",
		Lastname:  "Montenegro",
		Age:       35,
	},
	{
		Id:        "B002",
		Firstname: "Brian",
		Lastname:  "Montenegro",
		Age:       33,
	},
}

func getUsers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": users})
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userFound := funk.Find(users, func(u *dtos.User) bool { return u.Id == id })
	if userFound == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "User not exists"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": userFound})
}

func createUser(c *fiber.Ctx) error {
	var user *dtos.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	users = append(users, user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user.Id})
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user *dtos.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	userFound, ok := funk.Find(users, func(u *dtos.User) bool { return u.Id == id }).(*dtos.User)
	if !ok {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "User not exists"})
	}
	*userFound = *user

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": userFound.Id})
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if userFound := funk.Find(users, func(u *dtos.User) bool { return u.Id == id }); userFound == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "User not exists"})
	}

	users = funk.Filter(users, func(u *dtos.User) bool { return u.Id != id }).([]*dtos.User)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": id})
}
