package dto

type CategoryDto struct {
	Name string `json:"name" binding:"required"`
}
