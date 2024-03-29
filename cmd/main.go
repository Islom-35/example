package main

import (
	"example/pkg/db"
	"fmt"
	"log"
	"net/http"
	"os"

	"example/internal/adapters"
	"example/internal/app"
	"example/internal/controller/rest"
	"example/internal/domain"
)


// @title Example App API
// @version 1.0
// @description API Server for Example Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	db, err := db.New(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
	var product *domain.Product
	var users *domain.User
	db.AutoMigrate(&product)
	db.AutoMigrate(&users)

	defer db.Close()

	productRepo := adapters.NewProductRepository(db.DB)
	productService := app.NewProductService(productRepo)
	userRepo := adapters.NewUserRepository(db.DB)
	userService := app.NewUserService(userRepo)

	handlers := rest.NewHandler(productService, userService)

	// init & run server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v", os.Getenv("HTTP_PORT")),
		Handler: handlers.InitRouters(),
	}

	log.Println("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
