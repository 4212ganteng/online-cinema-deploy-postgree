package models

import "time"

type Transaction struct {
	ID     int          `json:"id"`
	UserID int          `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User   User         `json:"user"`
	FilmID int          `json:"film_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Film   FilmResponse `json:"film"`
	Status string       `json:"status"`
	Price int `json:"price"`
	TanggalOrder  time.Time `json:"tanggal_order" gorm:"default:Now()"`
}
type TransactionUserResponse struct {
	ID            int          `json:"id"`
	UserID        int          `json:"user_id"`
	User          User         `json:"user"`
	FilmID        int          `json:"film_id"`
	Film          FilmResponse `json:"film"`
	Status        string       `json:"status"`
	TanggalOrder  time.Time    `json:"tanggal_order" gorm:"default:Now()"`
}

func (TransactionUserResponse) TableName() string {
	return "transactions"
}
