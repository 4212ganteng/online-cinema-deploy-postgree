package repositories

import (
	"gorm.io/gorm"
	"online-cinema/models"
)

type CategoryRepository interface {
	GetCategory() ([]models.Categorie, error)
	GetCategoriID(ID int) (models.Categorie, error)
	CreateCategory(categories models.Categorie) (models.Categorie, error)
	UpdateCategorie(categories models.Categorie, ID int) (models.Categorie, error)
	DeleteCategory(categories models.Categorie, ID int) (models.Categorie, error)
}

func RepositoryCategies(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCategory() ([]models.Categorie, error) {
	var category []models.Categorie
	err := r.db.Find(&category).Error
	return category, err
}

func (r *repository) GetCategoriID(ID int) (models.Categorie, error) {
	var categorieId models.Categorie

	err := r.db.First(&categorieId, ID).Error
	return categorieId, err
}

func (r *repository) CreateCategory(category models.Categorie) (models.Categorie, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *repository) UpdateCategorie(category models.Categorie, ID int) (models.Categorie, error) {
	err := r.db.Model(&category).Where("id=?", ID).Updates(&category).Error
	return category, err
}

func (r *repository) DeleteCategory(category models.Categorie, ID int) (models.Categorie, error) {
	err := r.db.Delete(&category, ID).Error
	return category, err
}
