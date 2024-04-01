package main

import "fmt"

type passangerTrain struct {
	mediator mediator
}

func (g *passangerTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("PassangerTrain: Landing")
	} else {
		fmt.Println("PassangerTrain: Waiting")
	}
}

func (g *passangerTrain) departure() {
	fmt.Println("PassangerTrain: Leaving")
	g.mediator.notifyFree()
}

func (g *passangerTrain) permitArrival() {
	fmt.Println("PassangerTrain: Arrival Permitted. Landing")
}
