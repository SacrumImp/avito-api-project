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

	// Repositories
	userAccountRepo := repositories.NewUserAccountRepository(db)
	statusRepo := repositories.NewStatusRepository(db)
	houseRepo := repositories.NewHouseRepository(db)
	developerRepo := repositories.NewDeveloperRepository(db)
	flatRepo := repositories.NewFlatRepository(db)
	userTypeRepo := repositories.NewUserTypeRepository(db)

	// Services
	authService := services.NewAuthenticationService(userAccountRepo, userTypeRepo)
	flatService := services.NewFlatService(flatRepo, statusRepo)
	houseService := services.NewHouseService(houseRepo, developerRepo)

	// Handlers
	authHandler := handlers.NewAuthenticationHandler(authService)
	flatHandler := handlers.NewFlatHandler(flatService)
	houseHandler := handlers.NewHouseHandler(houseService, flatService)

	http.HandleFunc("/dummyLogin", authHandler.GetDummyJWT)
	http.HandleFunc("/register", authHandler.RegisterUser)
	http.Handle("/house/create",
		middleware.Authenticate(authService,
			middleware.RequireRoles([]string{string(models.Moderator)},
				http.HandlerFunc(houseHandler.CreateHouse))))
	http.Handle("/flat/create",
		middleware.Authenticate(authService,
			http.HandlerFunc(flatHandler.CreateFlat)))
	http.Handle("/flat/update",
		middleware.Authenticate(authService,
			middleware.RequireRoles([]string{string(models.Moderator)},
				http.HandlerFunc(flatHandler.UpdateFlatStatus))))
	http.Handle("/house/",
		middleware.Authenticate(authService,
			middleware.RequireRoles([]string{string(models.Moderator), string(models.Client)},
				http.HandlerFunc(houseHandler.GetFlatsByHouseID))))

	fmt.Println("API is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
