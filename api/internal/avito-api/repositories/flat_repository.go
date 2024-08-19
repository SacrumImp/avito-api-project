package repositories

import (
	"avito-api/internal/avito-api/models"
	"database/sql"
)

type FlatRepository interface {
	GetByHouseID(id int) ([]*models.Flat, error)
	CreateFlat(flat *models.Flat, statusId int) error
}

type SQLFlatRepository struct {
	DB *sql.DB
}

func NewFlatRepository(db *sql.DB) *SQLFlatRepository {
	return &SQLFlatRepository{DB: db}
}

func (repo *SQLFlatRepository) GetByHouseID(id int) ([]*models.Flat, error) {
	rows, err := repo.DB.Query(`
		SELECT
			number,
			house_id,
			price,
			number_of_rooms,
			status.title as status 
		FROM flat
		JOIN status on flat.status_id = status.id
		where house_id=1
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	flats := []*models.Flat{}
	for rows.Next() {
		flat := &models.Flat{}
		if err := rows.Scan(&flat.FlatId, &flat.HouseId, &flat.Price, &flat.Rooms, &flat.Status); err != nil {
			return nil, err
		}
		flats = append(flats, flat)
	}

	return flats, nil
}

func (repo *SQLFlatRepository) CreateFlat(flat *models.Flat, statusId int) error {
	query := `
		INSERT INTO flat (house_id, price, number_of_rooms, status_id)
		VALUES ($1, $2, $3, $4)
		RETURNING number;
	`
	if err := repo.DB.QueryRow(query, flat.HouseId, flat.Price, flat.Rooms, statusId).Scan(&flat.FlatId); err != nil {
		return err
	}
	return nil
}
