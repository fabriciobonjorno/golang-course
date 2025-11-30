package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.NoError(t, err)
	assert.Equal(t, 10.0, tax, "they should be equal")

	tax, err = CalculateTax(0.0)
	assert.Error(t, err)
	assert.Equal(t, 0.0, tax, "tax should be zero for negative amount")
	assert.Contains(t, err.Error(), "greater than zero")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.NotNil(t, err)
	assert.Equal(t, "error saving tax", err.Error())

	repository.AssertExpectations(t)

}
