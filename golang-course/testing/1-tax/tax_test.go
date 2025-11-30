package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amout := 500.0
	expectedTax := 5.0

	result := CalculateTax(amout)
	if result != expectedTax {
		t.Errorf("CalculateTax(%f) = %f; want %f", amout, result, expectedTax)
	}
}
