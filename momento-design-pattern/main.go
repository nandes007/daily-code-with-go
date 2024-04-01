package main

import "fmt"

func main() {
	caretaker := &caretaker{
		momentoArray: make([]*momento, 0),
	}
	originator := &originator{
		state: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMomento(originator.createMomento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())

	caretaker.addMomento(originator.createMomento())
	originator.setState("C")

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMomento(originator.createMomento())

	originator.restoreMomento(caretaker.getMomento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMomento(caretaker.getMomento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
