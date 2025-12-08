package database

import "apis/internal/entity"

// UserDB defines the interface for user database operations.
type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]*entity.Product, int64, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
