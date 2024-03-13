package main

import (
	"golangbyexample.com/encapsulation/model"
)

//Golang provides encapsulation at the package level. Go doesn't have any public,
//private or protected keyword. The only mechanism to control the visibility is using
//the capitalized and non-capitalized formats

//- Capitalized identifiers are exported. The capital letter indicates that this is an exported identifier.
//- Non-capitalized identifiers are not exported. The lowercase indicates that the identifier is not exported
//and will only be accessed from within the same package.

// There are five kinds of identifier which can be exported or non-exported
// 1. Structure (struct)
// 2. Structure's Method
// 3. Structure's Field
// 4. Function
// 5. Variable
func main() {
	model.Call()

	//Below function will raise error because call non-exported data within different package.
	// view.Call()
}
