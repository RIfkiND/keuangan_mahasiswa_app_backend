package repositories

import (
    "gorm.io/gorm"
    "keuangan/backend/internals/models"
)

type CategoryRepository struct {
    DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
    return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
    return r.DB.Create(category).Error
}

func (r *CategoryRepository) GetByID(id uint) (*models.Category, error) {
    var category models.Category
    err := r.DB.First(&category, id).Error
    return &category, err
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
    var categories []models.Category
    err := r.DB.Find(&categories).Error
    return categories, err
}

func (r *CategoryRepository) Update(category *models.Category) error {
    return r.DB.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
    return r.DB.Delete(&models.Category{}, id).Error
}