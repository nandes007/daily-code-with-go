package main

type taxSystem interface {
	calculateTax() int
}

type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	return i.income * i.taxPercentage / 100
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (s *singaporeTax) calculateTax() int {
	return s.income * s.taxPercentage / 100
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (u *usaTax) calculateTax() int {
	return u.income * u.taxPercentage / 100
}

func calculateTotalTax(taxSystem []taxSystem) int {
	totalTax := 0
	for _, t := range taxSystem {
		totalTax += t.calculateTax() //This is where runtime polymorphism happens
	}
	return totalTax
}
