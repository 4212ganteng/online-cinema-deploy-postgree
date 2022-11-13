package auth

type AuthResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
	Phone string `json:"phone"`
	Token    string `json:"token"`
}
type CheckAuthResponse struct {
	Id     int    `gorm:"type: int" json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Role string `gorm:"type: varchar(255)" json:"role"`
	Status string `gorm:"type: varchar(50)"  json:"status"`
}
