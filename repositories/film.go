package repositories

import (
	"gorm.io/gorm"
	"online-cinema/models"
)

type FilmRepository interface {
	GetFilm() ([]models.Film, error)
	GetfilmID(ID int) (models.Film, error)
	CreateFilm(film models.Film) (models.Film, error)
	UpdateFilm(film models.Film, ID int) (models.Film, error)
	DeleteFilm(film models.Film, ID int) (models.Film, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetFilm() ([]models.Film, error) {
	var film []models.Film
	err := r.db.Preload("Category").Find(&film).Error
	return film, err
}

func (r *repository) GetfilmID(ID int) (models.Film, error) {
	var filmId models.Film

	err := r.db.Preload("Category").First(&filmId, ID).Error
	return filmId, err
}

func (r *repository) CreateFilm(film models.Film) (models.Film, error) {
	err := r.db.Preload("Category").Create(&film).Error
	return film, err
}

func (r *repository) UpdateFilm(film models.Film, ID int) (models.Film, error) {
	err := r.db.Model(&film).Where("id=?", ID).Updates(&film).Error
	return film, err
}

func (r *repository) DeleteFilm(film models.Film, ID int) (models.Film, error) {
	err := r.db.Delete(&film, ID).Error
	return film, err
}
