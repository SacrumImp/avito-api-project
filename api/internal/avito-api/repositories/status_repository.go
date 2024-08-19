package repositories

import (
	"avito-api/internal/avito-api/models"
	"database/sql"
)

type StatusRepository interface {
	GetStatusByTitle(title string) (*models.Status, error)
}

type SQLStatusRepository struct {
	DB *sql.DB
}

func NewStatusRepository(db *sql.DB) *SQLStatusRepository {
	return &SQLStatusRepository{DB: db}
}

func (repo *SQLStatusRepository) GetStatusByTitle(title string) (*models.Status, error) {
	status := &models.Status{}

	query := `
		SELECT Id, Title FROM status where Title = $1
	`
	if err := repo.DB.QueryRow(query, title).Scan(&status.Id, &status.Title); err != nil {
		return nil, err
	}
	return status, nil
}
