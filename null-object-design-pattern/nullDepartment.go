package main

type nullDepartment struct {
	numberOfProfessors int
}

func (n *nullDepartment) getNumberOfProfessors() int {
	return n.numberOfProfessors
}

func (n *nullDepartment) getName() string {
	return "nullDepartment"
}
