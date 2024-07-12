package operation

func EarningBeforeTax(income, expenses float64) float64 {
	return income - expenses
}

func EarningAfterTax(earningBeforeTax, taxRate float64) float64 {
	tax := earningBeforeTax * (taxRate / 100)
	earningAfterTax := earningBeforeTax - tax
	return earningAfterTax
}
