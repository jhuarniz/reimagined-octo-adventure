package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jhuarniz/reimagined-octo-adventure/api/configs"
	"github.com/jhuarniz/reimagined-octo-adventure/api/controllers"
	"github.com/jhuarniz/reimagined-octo-adventure/api/models"
)

func main() {

	configs.DBConnection()
	configs.DB.AutoMigrate(&models.ImageRecognition{})

	// App Intance
	app := fiber.New()

	// Middlewares
	app.Use(logger.New())

	// Controllers
	controllers.IndexController(app)
	controllers.UserController(app)
	controllers.ImageRecognitionController(app)

	// Lauch
	app.Listen(":3000")
}
