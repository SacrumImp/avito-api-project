package repositories

import (
	"avito-api/internal/avito-api/models"
	"database/sql"
)

type UserAccountRepository interface {
	CreateUserAccount(userAccount *models.UserAccount, userTypeId int) error
	FindUserAccount(userAccount *models.UserAccount) error
}

type SQlUserAccountRepository struct {
	DB *sql.DB
}

func NewUserAccountRepository(db *sql.DB) *SQlUserAccountRepository {
	return &SQlUserAccountRepository{DB: db}
}

func (repo *SQlUserAccountRepository) CreateUserAccount(userAccount *models.UserAccount, userTypeId int) error {
	query := `
		with inserted_user as (
			INSERT INTO user_account (email, password_hash, user_type_id)
			VALUES ($1, $2, $3)
			RETURNING *
		)
		SELECT 
			inserted_user.id as id,
			user_type.Title as title
		FROM inserted_user
		JOIN user_type on user_type.id = inserted_user.user_type_id;
	`
	if err := repo.DB.QueryRow(query, userAccount.Email, userAccount.PasswordHash, userTypeId).Scan(&userAccount.UserId, &userAccount.UserType); err != nil {
		return err
	}
	return nil
}

func (repo *SQlUserAccountRepository) FindUserAccount(userAccount *models.UserAccount) error {
	query := `
		SELECT 
			user_type.Title as user_type,
			user_account.password_hash as password_hash
		FROM user_account
		JOIN user_type on user_type.id = user_account.user_type_id
		WHERE user_account.id=$1;
	`
	if err := repo.DB.QueryRow(query, userAccount.UserId).Scan(&userAccount.UserType, &userAccount.PasswordHash); err != nil {
		return err
	}
	return nil
}
