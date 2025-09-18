package repositories

import (
    "gorm.io/gorm"
    "gorestapi/internals/models"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *models.User) error {
    return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.DB.Where("email = ?", email).First(&user).Error
    return &user, err
}