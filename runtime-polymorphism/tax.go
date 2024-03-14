package main

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
