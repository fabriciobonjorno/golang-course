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

// go test -bench=. to run benchmark
func BenchmarkCalculateTax(b *testing.B) {
	amount := 500.0
	for i := 0; i < b.N; i++ {
		CalculateTax(amount)
	}
}

// go test -fuzz=.
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, 500.0, 750.0, 900.0, 800.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0 ", result)
		}
		if amount > 20000 && result != 20.0 {
			t.Errorf("Received %f but expected at least 1020.0 ", result)
		}
	})
}
