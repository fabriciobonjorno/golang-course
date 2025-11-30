package tax

import (
	"testing"
)

// generate coverage report
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out

func TestCalculateTax(t *testing.T) {
	amout := 500.0
	expectedTax := 10.0

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
		{amount: 100.0, expectedTax: 5.0},
		{amount: 250.0, expectedTax: 5},
		{amount: 500.0, expectedTax: 10.0},
		{amount: 800.0, expectedTax: 10.0},
	}

	for _, tc := range table {
		result := CalculateTax(tc.amount)
		if result != tc.expectedTax {
			t.Errorf("CalculateTax(%f) = %f; want %f", tc.amount, result, tc.expectedTax)
		}
	}
}
