package services

import (
	"gorestapi/internals/models"
	"gorestapi/internals/repositories"
)


type CategoryService struct {
	Repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		Repo: repo,
	}
}

func (s *CategoryService) Create(category *models.Category) error {
	return s.Repo.Create(category)
}

func (s *CategoryService) GetByID(id uint) (*models.Category, error) {
	return s.Repo.GetByID(id)
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.Repo.GetAll()
}

func (s *CategoryService) Update(category *models.Category) error {
	return s.Repo.Update(category)
}

func (s *CategoryService) Delete(id uint) error {
	return s.Repo.Delete(id)
}