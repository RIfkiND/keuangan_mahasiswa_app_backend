package services

import (
	"gorestapi/internals/models"
	"gorestapi/internals/repositories"
)

type PostService struct {
	Repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) Create(post *models.Post) error {
	return s.Repo.Create(post)
}

func (s *PostService) GetByID(id uint) (*models.Post, error) {
	return s.Repo.GetByID(id)
}

func (s *PostService) GetAll() ([]models.Post, error) {
	return s.Repo.GetAll()
}

func (s *PostService) Update(post *models.Post) error {
	return s.Repo.Update(post)
}

func (s *PostService) Delete(id uint) error {
	return s.Repo.Delete(id)
}