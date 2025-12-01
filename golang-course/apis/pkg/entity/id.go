package entity

import "github.com/google/uuid"

type ID = uuid.UUID

// NewID generates a new unique identifier. Using UUIDv7 for better sorting properties.
func NewID() ID {
	id, _ := uuid.NewV7()
	return ID(id)
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
