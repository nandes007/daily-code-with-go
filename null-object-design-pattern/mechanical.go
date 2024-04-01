package main

type mechanical struct {
	numberOfProfessors int
}

func (m *mechanical) getNumberOfProfessors() int {
	return m.numberOfProfessors
}

func (m *mechanical) getName() string {
	return "mechanical"
}
