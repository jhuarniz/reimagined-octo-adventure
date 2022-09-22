package models

type ImageRecognition struct {
	RequestId string `json:"requestId" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Filters   int    `json:"filters" validate:"required"`
	ImageUrl  string `json:"imageUrl" validate:"required"`
}
