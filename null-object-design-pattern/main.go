package main

import "fmt"

func main() {
	college1 := createCollege1()
	college2 := createCollege2()
	totalProfessors := 0
	departmentArray := []string{"computerscience", "mechanical", "civil", "electronics"}

	for _, departmentName := range departmentArray {
		d := college1.getDepartment(departmentName)
		totalProfessors += d.getNumberOfProfessors()
	}

	fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)

	//Reset the professors count
	totalProfessors = 0
	for _, departmentName := range departmentArray {
		d := college2.getDepartment(departmentName)
		totalProfessors += d.getNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}
