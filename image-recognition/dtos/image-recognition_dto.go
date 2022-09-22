package dtos

import "mime/multipart"

type ImageRecognition struct {
	Email     string                `json:"email" validate:"required,email"`
	Filters   string                `json:"filters" validate:"required"`
	ImageFile *multipart.FileHeader `json:"imageFile" validate:"required"`
}
