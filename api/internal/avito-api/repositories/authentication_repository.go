package repositories

import "database/sql"

type AuthenticationRepository interface {
}

type SQlAuthenticationRepository struct {
	DB *sql.DB
}

func NewAuthenticationRepository(db *sql.DB) *SQlAuthenticationRepository {
	return &SQlAuthenticationRepository{DB: db}
}
