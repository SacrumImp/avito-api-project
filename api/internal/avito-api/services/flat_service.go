package services

import (
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/repositories"
)

type FlatService struct {
	Repo repositories.FlatRepository
}

func NewFlatService(repo repositories.FlatRepository) *FlatService {
	return &FlatService{Repo: repo}
}

func (s *FlatService) GetByHouseID(id int) ([]*models.Flat, error) {
	flats, err := s.Repo.GetByHouseID(id)
	if err != nil {
		return nil, err
	}
	return flats, nil
}
