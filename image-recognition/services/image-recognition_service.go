package services

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jhuarniz/reimagined-octo-adventure/api/dtos"
)

var (
	validate = validator.New()
)

type imageRecognition interface {
	ImageRecognition(model dtos.ImageRecognition, c *fiber.Ctx) (string, error)
}

type imgreg struct{}

func NewImageRecognition() imageRecognition {
	return &imgreg{}
}

func (*imgreg) ImageRecognition(model dtos.ImageRecognition, c *fiber.Ctx) (string, error) {

	// validate
	err := validate.Struct(model)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return "", err
		}

		return "", err
	}

	// Save File
	c.SaveFile(model.ImageFile, fmt.Sprintf("./%s", model.ImageFile.Filename))

	return "123456", nil
}
