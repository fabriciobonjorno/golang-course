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
func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount      float64
		expectedTax float64
	}

	table := []calcTax{
		{amount: 100.0, expectedTax: 1.0},
		{amount: 250.0, expectedTax: 2.5},
		{amount: 400.0, expectedTax: 4.0},
		{amount: 800.0, expectedTax: 8.0},
	}

	for _, tc := range table {
		result := CalculateTax(tc.amount)
		if result != tc.expectedTax {
			t.Errorf("CalculateTax(%f) = %f; want %f", tc.amount, result, tc.expectedTax)
		}
	}
}
