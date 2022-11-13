package repositories

import (
	"online-cinema/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() ([]models.User, error)
	GetUserID(ID int) (models.User, error)
	// UpdateUser(user models.User, ID int) (models.User, error)
	// DeleteUser(user models.User, ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser() ([]models.User, error) {
	var user []models.User
	err := r.db.Preload("Transaction.Film").Preload("Transaction.User").Preload("Transaction.Film.Category").Find(&user).Error
	return user, err
}

func (r *repository) GetUserID(ID int) (models.User, error) {
	var UserId models.User

	err := r.db.Preload("Transaction.Film").Preload("Transaction.User").Preload("Transaction.Film.Category").First(&UserId, ID).Error
	return UserId, err
}

// func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
// 	err := r.db.Model(&user).Where("id=?", ID).Updates(&user).Error
// 	return user, err
// }

// func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
// 	err := r.db.Delete(&user, ID).Error
// 	return user, err
// }
