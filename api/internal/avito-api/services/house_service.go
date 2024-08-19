package services

import (
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/repositories"
)

type HouseService struct {
	HouseRepo     repositories.HouseRepository
	DeveloperRepo repositories.DeveloperRepository
}

func NewHouseService(houseRepo repositories.HouseRepository, developerRepo repositories.DeveloperRepository) *HouseService {
	return &HouseService{
		HouseRepo:     houseRepo,
		DeveloperRepo: developerRepo,
	}
}

func (s *HouseService) CreateHouse(houseInput *models.HouseInputObject) (*models.House, error) {
	house := &models.House{
		Address:   houseInput.Address,
		Year:      houseInput.Year,
		Developer: houseInput.Developer,
	}

	var developerId *int
	if houseInput.Developer != "" {
		developer, err := s.DeveloperRepo.GetDeveloperByTitle(houseInput.Developer)
		if err != nil {
			return nil, err
		}
		developerId = &developer.Id
	}

	house, err := s.HouseRepo.CreateHouse(house, developerId)
	if err != nil {
		return nil, err
	}
	return house, err
}
