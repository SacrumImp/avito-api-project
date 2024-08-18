package services

import (
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthenticationService struct {
	Repo repositories.AuthenticationRepository
}

func NewAuthenticationService(repo repositories.AuthenticationRepository) *AuthenticationService {
	return &AuthenticationService{Repo: repo}
}

var jwtKey = []byte("0b79d3d683a9f2eee06fafc2358aa25aa477c2339a63ae46e7b7092892f7eef9")

func generateJWT(role string) (string, error) {
	expirationTime := time.Now().Add(8 * time.Hour)

	claims := &models.Claim{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func (s *AuthenticationService) GetDummyJWT(userType models.UserTypesEnum) (*models.Token, error) {
	tokenString, err := generateJWT(string(userType))
	if err != nil {
		return nil, err
	}
	token := &models.Token{
		Token: tokenString,
	}
	return token, nil
}
