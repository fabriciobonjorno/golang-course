package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Test Product", 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, "Test Product", p.Name)
	assert.Equal(t, 100.0, p.Price)
	assert.NotEmpty(t, p.ID)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 100.0)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Test Product", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Test Product", -10.0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsInvalid, err)
}
