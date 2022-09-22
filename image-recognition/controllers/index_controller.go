package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jhuarniz/reimagined-octo-adventure/api/configs"
)

func IndexController(router fiber.Router) {
	indexGroup := router.Group("/")
	indexGroup.Get("", index)
}

func index(c *fiber.Ctx) error {
	myvar := configs.EnvS3APIKey()
	return c.SendString("Hello World!" + myvar)
}
