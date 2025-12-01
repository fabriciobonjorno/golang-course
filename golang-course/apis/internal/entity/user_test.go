package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@john.com", "securepassword")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@john.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Jane Doe", "jane@jane.com", "securepassword")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("securepassword"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.NotEqual(t, "securepassword", user.Password) // Ensure password is hashed
}
