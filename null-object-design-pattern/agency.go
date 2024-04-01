package main

func createCollege1() *college {
	college := &college{}
	college.addDepartment("computerscience", 4)
	college.addDepartment("mechanical", 5)
	return college
}

func createCollege2() *college {
	college := &college{}
	college.addDepartment("computerscience", 2)
	return college
}
