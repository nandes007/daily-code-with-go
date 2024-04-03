package main

import "fmt"

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	//Calculate area for square. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(c *circle) {
	//Calculate area for circle. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) visitForrectangle(r *rectangle) {
	//Calculate area for rectangle. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for reactangle")
}
