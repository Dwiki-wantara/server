package repositories

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategori() ([]models.Category, error)
	GetCategory(ID int) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(category models.Category) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategori() ([]models.Category, error) {
	var categori []models.Category
	err := r.db.Find(&categori).Error

	return categori, err
}

func (r *repository) GetCategory(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, ID).Error

	return category, err
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error

	return category, err
}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error

	return category, err
}

func (r *repository) DeleteCategory(category models.Category) (models.Category, error) {
	err := r.db.Delete(&category).Error

	return category, err
}
