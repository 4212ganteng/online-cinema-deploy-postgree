package auth

type AuthRequest struct {
	Email    string `gorm: "type : varchar(255)" form :"email" json:"email"`
	Password string `gorm: "type : varchar(255)" form : "password" json:"password" `
	FullName string `gorm: "type : varchar(255)" from :"fullName" json:"fullName"`
	Phone    string `gorm: "type : varchar(255)" from :"phone" json:"phone"`
}
