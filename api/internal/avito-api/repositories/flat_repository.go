package repositories

import (
	"avito-api/internal/avito-api/models"
	"database/sql"
)

type FlatRepository interface {
	GetByHouseID(id int) ([]*models.Flat, error)
	CreateFlat(flat *models.Flat, statusId int) error
	UpdateFlatStatus(flat *models.Flat, statusId int) error
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
		WITH inserted_flat as (
			INSERT INTO flat (house_id, price, number_of_rooms, status_id)
			VALUES ($1, $2, $3, $4)
			RETURNING *
		)
		SELECT 
			inserted_flat.number as number,
			status.Title as status
		FROM inserted_flat
		JOIN status on inserted_flat.status_id = status.id;
	`
	if err := repo.DB.QueryRow(query, flat.HouseId, flat.Price, flat.Rooms, statusId).Scan(&flat.FlatId, &flat.Status); err != nil {
		return err
	}
	return nil
}

func (repo *SQLFlatRepository) UpdateFlatStatus(flat *models.Flat, statusId int) error {
	query := `
		WITH updated_flat as (
			UPDATE flat 
			SET status_id = $3
			WHERE house_id = $1 and number = $2
			RETURNING *
		)
		SELECT 
			updated_flat.price as price,
			updated_flat.number_of_rooms as rooms,
			status.Title as status
		FROM updated_flat
		JOIN status on updated_flat.status_id = status.id;
	`
	if err := repo.DB.QueryRow(query, flat.HouseId, flat.FlatId, statusId).Scan(&flat.Price, &flat.Rooms, &flat.Status); err != nil {
		return err
	}
	return nil
}
