package services

import (
    "mime/multipart"
    "gorestapi/internals/repositories"
)

type PostImageService struct {
    Repo *repositories.PostImageRepository
}

func NewPostImageService(repo *repositories.PostImageRepository) *PostImageService {
    return &PostImageService{Repo: repo}
}

func (s *PostImageService) UploadImage(file multipart.File, fileName string) (string, error) {
    return s.Repo.UploadImage(file, fileName)
}