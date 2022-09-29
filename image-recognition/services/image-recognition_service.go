package services

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jhuarniz/reimagined-octo-adventure/api/configs"
	"github.com/jhuarniz/reimagined-octo-adventure/api/dtos"
	"github.com/jhuarniz/reimagined-octo-adventure/api/models"
	"github.com/jinzhu/copier"
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
		return "", err
	}

	// Save Image File
	c.SaveFile(model.ImageFile, fmt.Sprintf("./%s", model.ImageFile.Filename))

	// Parse from dto to model
	var newModel models.ImageRecognition
	copier.Copy(&newModel, &model)

	// Persistence in DataBase
	createdUser := configs.DB.Create(&newModel)
	err = createdUser.Error
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	// Return ID
	return newModel.ID.String(), nil
}
