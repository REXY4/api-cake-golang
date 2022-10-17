package dto

import (
	"mime/multipart"
)

type CakeCreateDTO struct {
	ID          uint64                `json:"id" form:"id"`
	Title       string                `json:"title" form:"title" binding:"required"`
	Description string                `json:"description" form:"description" binding:"required"`
	Rating      string                `json:"rating"  form:"rating" binding:"required"`
	Image       string                `json:"image" form:"image" `
	ImageFile   *multipart.FileHeader `json:"imageFile" form:"imageFile" binding:"required"`
}
