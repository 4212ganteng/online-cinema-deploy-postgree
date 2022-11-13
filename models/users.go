package models

type User struct {
	ID          int                       `json:"id"`
	Email       string                    `json:"email"`
	FullName    string                    `json:"fullName"`
	Password    string                    `json:"password"`
	Phone       string                    `json:"phone"`
	Role        string                    `json:"role"`
	Transaction []TransactionUserResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaction"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

type UsersProfileResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (UserResponse) TableName() string {
	return "users"
}


func (UsersProfileResponse) TableName() string {
	return "users"
}
