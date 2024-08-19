package repositories

import (
	"avito-api/internal/avito-api/models"
	"database/sql"
	"time"
)

type HouseRepository interface {
	CreateHouse(house *models.House, developerId *int) error
}

type SQLHouseRepository struct {
	DB *sql.DB
}

func NewHouseRepository(db *sql.DB) *SQLHouseRepository {
	return &SQLHouseRepository{DB: db}
}

func (repo *SQLHouseRepository) CreateHouse(house *models.House, developerId *int) error {
	var developerIdParam sql.NullInt32
	if developerId != nil {
		developerIdParam = sql.NullInt32{Int32: int32(*developerId), Valid: true}
	} else {
		developerIdParam = sql.NullInt32{Valid: false}
	}

	var insertedAt sql.NullTime
	var lastFlatAddedAt sql.NullTime
	query := `
		INSERT INTO house (address, year_of_construction, developer_id)
		VALUES ($1, $2, $3)
		RETURNING id, inserted_at, last_flat_added_at;
	`
	if err := repo.DB.QueryRow(query, house.Address, house.Year, developerIdParam).Scan(&house.HouseId, &insertedAt, &lastFlatAddedAt); err != nil {
		return err
	}

	if insertedAt.Valid {
		house.CreatedAt = insertedAt.Time.Format(time.RFC3339)
	}
	if lastFlatAddedAt.Valid {
		house.UpdateAt = lastFlatAddedAt.Time.Format(time.RFC3339)
	}

	return nil
}
