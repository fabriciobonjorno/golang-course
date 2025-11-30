package tax

import (
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
