package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jhuarniz/reimagined-octo-adventure/api/dtos"
	"github.com/jhuarniz/reimagined-octo-adventure/api/services"
)

func ImageRecognitionController(router fiber.Router) {
	recognitionGroup := router.Group("/image-recognition")
	recognitionGroup.Post("/v1/load", imageRecognition)
}

func imageRecognition(c *fiber.Ctx) error {
	log.Println("imageRecognition - Start")

	// Get File
	imageFile, err := c.FormFile("fileUpload")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Call to Service
	imgrec := dtos.ImageRecognition{c.FormValue("email"), c.FormValue("filters"), imageFile}
	requestID, err := services.NewImageRecognition().ImageRecognition(imgrec, c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			dtos.ResponseDto{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"error": err.Error()},
			})
	}

	// All OK
	log.Println("imageRecognition - End")
	return c.Status(fiber.StatusAccepted).JSON(
		dtos.ResponseDto{
			StatusCode: fiber.StatusAccepted,
			Message:    "success",
			Data:       &fiber.Map{"requestID": requestID},
		})
}
