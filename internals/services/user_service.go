package services

import (
    "gorestapi/internals/models"
    "gorestapi/internals/repositories"
)

type UserService struct {
    Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) Register(user *models.User) error {
    return s.Repo.Create(user)
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
    return s.Repo.FindByEmail(email)
}