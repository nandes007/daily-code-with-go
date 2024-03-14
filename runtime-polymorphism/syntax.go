package main

import "fmt"

// RunTime polymorphism means that a call is resolved at runtime. It is achieved in GO using interfaces.
func call() {
	indianTax := &indianTax{
		taxPercentage: 30,
		income:        1000,
	}

	singaporeTax := &singaporeTax{
		taxPercentage: 10,
		income:        2000,
	}

	usaTax := &usaTax{
		taxPercentage: 40,
		income:        500,
	}

	taxSystems := []taxSystem{indianTax, singaporeTax, usaTax}
	totalTax := calculateTotalTax(taxSystems)

	fmt.Printf("Total Tax is %d\n", totalTax)
}

func calculateTotalTax(taxSystems []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystems {
		totalTax += t.calculateTax() //This is where runtime polymorphism happens
	}
	return totalTax
}
