package dtos

import "github.com/gofiber/fiber/v2"

type ResponseDto struct {
	StatusCode int        `json:"statusCode"`
	Message    string     `json:"message"`
	Data       *fiber.Map `json:"data"`
}
