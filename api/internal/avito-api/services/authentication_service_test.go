package services

import (
	"avito-api/internal/avito-api/models"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var mockJwtKey = []byte("0b79d3d683a9f2eee06fafc2358aa25aa477c2339a63ae46e7b7092892f7eef9")

func TestGenerateJWT(t *testing.T) {
	role := "client"

	originalJwtKey := jwtKey
	defer func() { jwtKey = originalJwtKey }()
	jwtKey = mockJwtKey

	tokenString, err := generateJWT(role)
	assert.NoError(t, err, "Expected no error, but got one")
	assert.NotEmpty(t, tokenString, "Expected a token, but got an empty string")
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return mockJwtKey, nil
	})
	assert.NoError(t, err, "Failed to parse token")
	if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
		assert.Equal(t, role, claims.Role, "Expected role to be %s, but got %s", role, claims.Role)
	} else {
		t.Fatalf("Token claims are invalid")
	}
}

func TestGenerateJWTError(t *testing.T) {
	originalJwtKey := jwtKey
	defer func() { jwtKey = originalJwtKey }()

	invalidJwtKey := []byte{}
	jwtKey = invalidJwtKey

	tokenString, err := generateJWT("client")
	assert.Error(t, err, "Expected an error, but got none")
	assert.Empty(t, tokenString, "Expected an empty token string, but got one")
}
