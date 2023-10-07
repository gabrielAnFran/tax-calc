package taxcalc

import "testing"

func TestCalculateTax(t *testing.T) {
	type value_with_expected struct {
		income   float64
		expected float64
		state    string
	}
	values := []value_with_expected{
		{income: 10, expected: 10.0 * (10.0 / 100.0), state: "RS"},
		{income: 10, expected: 10.0*(10.0/100.0) + 10, state: "SP"},
		{income: 100, expected: 100.0 * (10.0 / 100.0), state: "RJ"},
		{income: 1000, expected: 1000.0 * (20.0 / 100.0), state: "MA"},
	}

	for _, value := range values {
		actual := CalculateTax(value.income, value.state)
		if actual != value.expected {
			t.Errorf("CalculateTax(%f): expected %f, actual %f", value.income, value.expected, actual)
		}
	}
}
