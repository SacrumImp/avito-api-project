package services

import (
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/repositories"
)

type FlatService struct {
	FlatRepo   repositories.FlatRepository
	StatusRepo repositories.StatusRepository
}

func NewFlatService(flatRepo repositories.FlatRepository, statusRepo repositories.StatusRepository) *FlatService {
	return &FlatService{
		FlatRepo:   flatRepo,
		StatusRepo: statusRepo,
	}
}

func (s *FlatService) GetByHouseID(id int) ([]*models.Flat, error) {
	flats, err := s.FlatRepo.GetByHouseID(id)
	if err != nil {
		return nil, err
	}
	return flats, nil
}

func (s *FlatService) CreateFlat(flatInput *models.FlatInputObject) (*models.Flat, error) {
	flat := &models.Flat{
		HouseId: flatInput.HouseId,
		Price:   flatInput.Price,
		Rooms:   flatInput.Rooms,
	}

	status, err := s.StatusRepo.GetStatusByTitle(models.Created)
	if err != nil {
		return nil, err
	}

	if err := s.FlatRepo.CreateFlat(flat, status.Id); err != nil {
		return nil, err
	}
	return flat, nil
}

func (s *FlatService) UpdateFlatStatus(flatUpdate *models.FlatUpdateObject) (*models.Flat, error) {
	flat := &models.Flat{
		HouseId: flatUpdate.HouseId,
		FlatId:  flatUpdate.FlatId,
	}

	status, err := s.StatusRepo.GetStatusByTitle(flatUpdate.Status)
	if err != nil {
		return nil, err
	}

	if err := s.FlatRepo.UpdateFlatStatus(flat, status.Id); err != nil {
		return nil, err
	}
	return flat, nil
}

func (s *FlatService) FilterByRole(flats []*models.Flat, role string) ([]*models.Flat, error) {

	if role == string(models.Moderator) {
		return flats, nil
	}

	filteredFlats := []*models.Flat{}
	for _, flat := range flats {
		if flat.Status == models.Approved {
			filteredFlats = append(filteredFlats, flat)
		}
	}
	return filteredFlats, nil
}
