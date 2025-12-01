package database

import "apis/internal/entity"

// UserDB defines the interface for user database operations.
type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
