package main

import (
	"avito-api/internal/avito-api/handlers"
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

	flatRepo := repositories.NewFlatRepository(db)
	flatService := services.NewFlatService(flatRepo)
	flatHandler := handlers.NewFlatHandler(flatService)

	http.HandleFunc("/house/", flatHandler.GetByHouseID)

	fmt.Println("API is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
