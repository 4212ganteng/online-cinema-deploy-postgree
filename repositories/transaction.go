package repositories

import (
	"online-cinema/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	GetTransactionByUserID(userID int) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetTransactionID(ID int) (models.Transaction, error)
	// Declare GetOneTransaction repository method here ...
	GetOneTransaction(ID string) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	// Declare UpdateTransaction repository method here ...
	UpdateTransaction(status string, ID string) error
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Film.Category").Preload("User").Find(&transaction).Error
	return transaction, err
}

func (r *repository) GetTransactionByUserID(userID int) (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.Preload("Film.Category").Preload("User").First(&transactions, "user_id = ?", userID).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.Preload("Film").First(&transactions, "id = ?", ID).Error

	return transactions, err
}
func (r *repository) GetTransactionID(ID int) (models.Transaction, error) {
	var transactionId models.Transaction

	err := r.db.Preload("Film").First(&transactionId, ID).Error
	return transactionId, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

// Create GetOneTransaction method here ...
func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Film").Preload("Film.User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

// Create UpdateTransaction method here ...
func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("Film").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var film models.Film
		r.db.First(&film, transaction.Film.ID)

		r.db.Save(&film)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}
