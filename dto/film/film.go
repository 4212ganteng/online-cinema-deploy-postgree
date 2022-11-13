package film

type CreateFilmRequest struct {
	Title       string `json:"title" form:"title" gorm:"type: varchar(255)"`
	CategoryID  int    `json:"categorie_id" form:"category_id" gorm:"type: int"`
	Description string `json:"description" gorm:"type:text" form:"description"`
	Price       int    `json:"price" form:"price" gorm:"type: int"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
	FilmUrl     string `json:"filmUrl" form:"filmUrl" gorm:"type: varchar(255)"`
}
