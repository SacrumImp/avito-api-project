package main

import (
	"avito-api/internal/avito-api/handlers"
	"avito-api/internal/avito-api/middleware"
	"avito-api/internal/avito-api/models"
	"avito-api/internal/avito-api/repositories"
	"avito-api/internal/avito-api/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPasssword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPasssword, dbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	authRepo := repositories.NewAuthenticationRepository(db)
	authService := services.NewAuthenticationService(authRepo)
	authHandler := handlers.NewAuthenticationHandler(authService)

	flatRepo := repositories.NewFlatRepository(db)
	flatService := services.NewFlatService(flatRepo)
	flatHandler := handlers.NewFlatHandler(flatService)

	houseRepo := repositories.NewHouseRepository(db)
	developerRepo := repositories.NewDeveloperRepository(db)
	houseService := services.NewHouseService(houseRepo, developerRepo)
	houseHandler := handlers.NewHouseHandler(houseService)

	http.HandleFunc("/dummyLogin", authHandler.GetDummyJWT)
	http.Handle("/house/create",
		middleware.Authenticate(authService,
			middleware.RequireRole(string(models.Moderator),
				http.HandlerFunc(houseHandler.CreateHouse))))
	http.HandleFunc("/house/", flatHandler.GetByHouseID)

	fmt.Println("API is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
