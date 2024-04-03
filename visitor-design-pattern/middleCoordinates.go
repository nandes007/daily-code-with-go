package main

import "fmt"

type middleCoordinates struct {
	x int
	y int
}

func (m *middleCoordinates) visitForSquare(s *square) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (m *middleCoordinates) visitForCircle(c *circle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for circle")
}

func (m *middleCoordinates) visitForrectangle(r *rectangle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for rectangle")
}
