package main

type college struct {
	departments []department
}

func (c *college) addDepartment(departmentName string, numOfProfessors int) {
	if departmentName == "computerscience" {
		computerScienceDepartment := &computerscience{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, computerScienceDepartment)
	}

	if departmentName == "mechanical" {
		mechanicalDepartment := &mechanical{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, mechanicalDepartment)
	}
}

func (c *college) getDepartment(departmentName string) department {
	for _, department := range c.departments {
		if department.getName() == departmentName {
			return department
		}
	}

	return &nullDepartment{}
}
