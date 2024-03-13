package view

// This function to call data for exported and non-exported data within same package.
// Below function will raise error, because call non-exported data within different package.
// func Call() {
// 	p := &model.Person{
// 		Name: "test",
// 		age:  21,
// 	}
// 	fmt.Println(p)
// 	c := &model.company{}
// 	fmt.Println(c)

// 	//Structure's Method
// 	fmt.Println(p.Name)
// 	fmt.Println(p.age)

// 	//Function
// 	person2 := model.GetPerson()
// 	fmt.Println(person2)
// 	companyName := model.getCompanyName()
// 	fmt.Println(companyName)

// 	//Variables
// 	fmt.Println(model.companyLocation)
// 	fmt.Println(model.CompanyName)
// }
