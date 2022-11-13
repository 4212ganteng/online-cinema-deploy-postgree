package models

type Film struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Category    Categorie `json:"categorie"`
	CategoryID  int       `json:"categorie_id" form:"categorie_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Price       int       `json:"price"`
	FilmUrl     string    `json:"filmUrl"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
}

type FilmResponse struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Category    CategoryResponse `json:"categorie"`
	CategoryID  int              `json:"category_id" form:"category_id"`
	Price       int              `json:"price"`
	FilmUrl     string           `json:"filmUrl"`
	Description string           `json:"description"`
	Image       string           `json:"image"`
}

func (FilmResponse) TableName() string {
	return "films"
}
