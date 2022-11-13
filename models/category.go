package models

type Categorie struct {
	ID   int    `json:"id" `
	Name string `json:"name"`
}
type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CategoryResponse) TableName() string {
	return "categories"
}
