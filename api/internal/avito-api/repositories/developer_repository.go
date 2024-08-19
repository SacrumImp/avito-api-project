package repositories

import (
	"avito-api/internal/avito-api/models"
	"database/sql"
)

type DeveloperRepository interface {
	GetDeveloperByTitle(title string) (*models.Developer, error)
}

type SQLDeveloperRepository struct {
	DB *sql.DB
}

func NewDeveloperRepository(db *sql.DB) *SQLDeveloperRepository {
	return &SQLDeveloperRepository{DB: db}
}

func (repo *SQLDeveloperRepository) GetDeveloperByTitle(title string) (*models.Developer, error) {
	developer := &models.Developer{}

	query := `
		SELECT Id, Title FROM developer where Title = $1
	`
	if err := repo.DB.QueryRow(query, title).Scan(&developer.Id, &developer.Title); err != nil {
		return nil, err
	}
	return developer, nil
}
