package dto

type ReceipeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	CategoryId  uint   `json:"category_id"`
	Category    string `json:"category"`
	Time        string `json:"time"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Date        string `json:"date"`
}

type ReceipeResponses []ReceipeResponse
