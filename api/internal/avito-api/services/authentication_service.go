package services

import (
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/repositories"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService struct {
	UserAccountRepo repositories.UserAccountRepository
	UserTypeRepo    repositories.UserTypeRepository
}

func NewAuthenticationService(userAccountRepo repositories.UserAccountRepository, userTypeRepo repositories.UserTypeRepository) *AuthenticationService {
	return &AuthenticationService{
		UserAccountRepo: userAccountRepo,
		UserTypeRepo:    userTypeRepo,
	}
}

var jwtKey = []byte("0b79d3d683a9f2eee06fafc2358aa25aa477c2339a63ae46e7b7092892f7eef9")

func generateJWT(role string) (string, error) {
	if len(jwtKey) == 0 {
		return "", fmt.Errorf("invalid signing key")
	}

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

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func checkPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
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

func (s *AuthenticationService) DecodeJWT(tokenStr string) (*models.Claim, error) {
	claims := &models.Claim{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func (s *AuthenticationService) CreateUser(registerUser *models.UserRegisterObject) (*models.UserLogin, error) {
	hashedPassword, err := hashPassword(registerUser.Password)
	if err != nil {
		return nil, err
	}

	userAccount := &models.UserAccount{
		Email:        registerUser.Email,
		PasswordHash: hashedPassword,
	}

	userType, err := s.UserTypeRepo.GetUserTypeByTitle(registerUser.UserType)
	if err != nil {
		return nil, err
	}

	if err := s.UserAccountRepo.CreateUserAccount(userAccount, userType.Id); err != nil {
		return nil, err
	}

	userLogin := &models.UserLogin{
		UserId: userAccount.UserId,
	}

	return userLogin, nil
}

func (s *AuthenticationService) LoginUser(userLogin *models.UserLoginObject) (*models.Token, error) {
	userAccount := &models.UserAccount{
		UserId: userLogin.UserId,
	}

	if err := s.UserAccountRepo.FindUserAccount(userAccount); err != nil {
		return nil, err
	}

	if err := checkPasswordHash(userAccount.PasswordHash, userLogin.Password); err != nil {
		return nil, err
	}

	tokenString, err := generateJWT(userAccount.UserType)
	if err != nil {
		return nil, err
	}
	token := &models.Token{
		Token: tokenString,
	}

	return token, nil
}
