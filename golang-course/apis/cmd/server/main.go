package main

import (
	"apis/configs"
	"apis/internal/entity"
	"apis/internal/infra/database"
	"apis/internal/infra/webserver/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	// Initialize Product database and handler here
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB) // Create ProductHandler

	// using Chi router or any other router of your choice
	r := chi.NewRouter()

	// logs all requests
	r.Use(middleware.Logger)

	// Register the CreateProduct route
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)

	http.ListenAndServe(":8000", r)

	// Register the CreateProduct route
	// using default mux for simplicity
	// http.HandleFunc("/products", productHandler.CreateProduct)

	// http.ListenAndServe(":8000", nil)
}
