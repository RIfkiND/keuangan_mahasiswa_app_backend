package repositories

import (
    "gorm.io/gorm"
    "gorestapi/internals/models"
)

type PostRepository struct {
    DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
    return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *models.Post) error {
    return r.DB.Create(post).Error
}

func (r *PostRepository) GetByID(id uint) (*models.Post, error) {
    var post models.Post
    err := r.DB.First(&post, id).Error
    return &post, err
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
    var posts []models.Post
    err := r.DB.Find(&posts).Error
    return posts, err
}

func (r *PostRepository) Update(post *models.Post) error {
    return r.DB.Save(post).Error
}

func (r *PostRepository) Delete(id uint) error {
    return r.DB.Delete(&models.Post{}, id).Error
}