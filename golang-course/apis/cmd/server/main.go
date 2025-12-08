package main

import (
	"apis/configs"
	"apis/internal/entity"
	"apis/internal/infra/database"
	"apis/internal/infra/webserver/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandlers(userDB, configs.TokenAuth, configs.JWTExpiresIn) // add jwt an expiration on user EP

	// using Chi router or any other router of your choice
	r := chi.NewRouter()

	// logs all requests
	r.Use(middleware.Logger)

	// Products routes
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)

	// Register the CreateProduct route
	// using default mux for simplicity
	// http.HandleFunc("/products", productHandler.CreateProduct)

	// http.ListenAndServe(":8000", nil)
}
