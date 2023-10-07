package taxcalc

func CalculateTax(income float64, state string) float64 {
	var tax float64
	if income > 100 {
		tax = income * 20 / 100
	} else {
		tax = income * 10 / 100
	}

	if state == "SP" {
		tax = tax + 12
	}
	return tax
}