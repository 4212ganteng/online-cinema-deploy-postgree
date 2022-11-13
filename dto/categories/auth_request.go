package categories

type CategoryRequest struct {
	Name string `gorm: "type : varchar(255)" from :"name" json:"name"`
}
