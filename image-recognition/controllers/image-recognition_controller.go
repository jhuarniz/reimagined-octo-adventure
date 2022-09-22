package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jhuarniz/reimagined-octo-adventure/api/dtos"
	"github.com/jhuarniz/reimagined-octo-adventure/api/services"
)

func ImageRecognitionController(router fiber.Router) {
	recognitionGroup := router.Group("/image-recognition")
	recognitionGroup.Post("/v1/load", imageRecognition)
}

func imageRecognition(c *fiber.Ctx) error {
	fmt.Println("imageRecognition - Start")

	// Get Email
	email := c.FormValue("email")
	fmt.Println(email)

	// Get Filter
	filters := c.FormValue("filters")
	fmt.Println(filters)

	// Get File
	imageFile, err := c.FormFile("fileUpload")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Println(imageFile.Filename, imageFile.Size, imageFile.Header["Content-Type"][0])

	// Call to Service
	imgrec := dtos.ImageRecognition{
		Email:     email,
		Filters:   filters,
		ImageFile: imageFile,
	}
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
	return c.Status(fiber.StatusAccepted).JSON(
		dtos.ResponseDto{
			StatusCode: fiber.StatusAccepted,
			Message:    "success",
			Data:       &fiber.Map{"requestID": requestID},
		})
}
