package repositories

import (
	"avito-api/internal/avito-api/models"
	"database/sql"
)

type UserTypeRepository interface {
	GetUserTypeByTitle(title string) (*models.UserType, error)
}

type SQLUserTypeRepository struct {
	DB *sql.DB
}

func NewUserTypeRepository(db *sql.DB) *SQLUserTypeRepository {
	return &SQLUserTypeRepository{DB: db}
}

func (repo *SQLUserTypeRepository) GetUserTypeByTitle(title string) (*models.UserType, error) {
	userType := &models.UserType{}

	query := `
		SELECT Id, Title FROM user_type where Title = $1
	`
	if err := repo.DB.QueryRow(query, title).Scan(&userType.Id, &userType.Title); err != nil {
		return nil, err
	}
	return userType, nil
}
