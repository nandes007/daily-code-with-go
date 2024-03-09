package main

import "fmt"

// Declaring map
func delaringMap() {
	//Declare
	//Using map literal
	fmt.Println("Declare Map")
	employeeSalary := map[string]int{}
	fmt.Println(employeeSalary)

	//Initialize using map literal
	employeeSalary = map[string]int{
		"John": 1000,
		"Sam":  1200,
	}

	//Adding a key value
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)
}

// Declaring using var keyword
// One use case of having a map declared with var keyword is when an already existing map
// needs to be assigned to it or when we want to assign the result of a function
func delareUsingVar() {
	var employeeSalary map[string]int
	fmt.Println(employeeSalary)
	// employeeSalary["Tom"] = 2000 // This will run it to panic as the map nil
}

// Declaring using make
func declaringUsingMake() {
	//Declare
	employeeSalary := make(map[string]int)
	//Adding a key value
	fmt.Println("\nDeclaring map using make")
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)
}

//Map Operations
//- Add a key-value pair
//- Update a key
//- Get the value corresponding to a key
//- Delete a key-value pair
//- Check if key exists

func addingKeyValuePair() {
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	fmt.Println("\nAdding a key value")
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)
}

func updateMap() {
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	fmt.Println("\nUpdate Map")
	fmt.Println("Before update")
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)

	fmt.Println("After update")
	employeeSalary["Tom"] = 3000
	fmt.Println(employeeSalary)
}

func retrieveValueCorrespodingToKey() {
	fmt.Println("\nRetrieve Value Map Corresponding to key")
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	employeeSalary["Tom"] = 2000

	//Retrieve a value
	salary := employeeSalary["Tom"]
	fmt.Println(salary)
	fmt.Printf("Salary: %d", salary)
}

func deleteKeyValuePair() {
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	fmt.Println("\n\nDelete Key Value Pair")
	fmt.Println("Adding key")
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)

	fmt.Println("\nDeleting key")
	delete(employeeSalary, "Tom")
	fmt.Println(employeeSalary)
}

func checkIfKeyExists() {
	//Signature
	//val, ok := mapName[key]
	//There two cases
	//If the key exists val variable be the value of the key in the map and ok variable will be true
	//If the key doesn't exist val variable will be default zero value of value type and ok will be false.
	fmt.Println("\nCheck if key exist")
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	employeeSalary["Tom"] = 1200
	fmt.Println("Key exist case")
	val, ok := employeeSalary["Tom"]
	fmt.Printf("Val: %d, ok: %t\n", val, ok)

	fmt.Println("\nKey doesn't exists case")
	val, ok = employeeSalary["Sam"]
	fmt.Printf("Val: %d, ok: %t\n", val, ok)

	//NOTE
	//In case we only want to check if a key is present and val is no needed, then blank identifier ie:
	// _, ok := employeeSalary["Tom"]
}

// Function on map
// - len() function
func functionOnMap() {
	fmt.Println("\nFunction of map")
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	employeeSalary["Tom"] = 2000
	employeeSalary["Sam"] = 1200

	lenOfMap := len(employeeSalary)
	fmt.Println(lenOfMap)
}

func zeroValue() {
	fmt.Println("\nZero value")
	var employeeSalary map[string]int
	if employeeSalary == nil {
		fmt.Println("employeeSalary map is nil")
	}
}

func mapAreReferencedDataTypes() {
	fmt.Println("\nMap Are Referenced Data Type")
	// Map are reference data types.
	// So on assigning one map to new variable, then both both variable refers to same map.
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	employeeSalary["Tom"] = 2000
	employeeSalary["Sam"] = 1200

	eS := employeeSalary

	//Change employeeSalary
	employeeSalary["John"] = 3000
	fmt.Println("Changing employeeSalary Map")
	fmt.Printf("employeeSalary: %v\n", employeeSalary)
	fmt.Printf("eS: %v\n", eS)

	//Change eS
	eS["John"] = 4000
	fmt.Println("\nChanging eS Map")
	fmt.Printf("employeeSalary: %v\n", employeeSalary)
	fmt.Printf("eS: %v\n", eS)
}

// Iterate over a map
func iterateMap() {
	fmt.Println("\nIterate a Map")
	//Declare
	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	fmt.Println("\nIterating over keys and values")
	for k, v := range sample {
		fmt.Printf("key: %s value: %s\n", k, v)
	}

	fmt.Println("\nIteration over only keys")
	for k := range sample {
		fmt.Printf("key: %s\n", k)
	}

	fmt.Println("\nIteration over only values")
	for _, v := range sample {
		fmt.Printf("value: %s\n", v)
	}

	fmt.Println("\nGet list of keys")
	listOfKeys := getAllKeys(sample)
	fmt.Println(listOfKeys)
}

// Maps are not safe for concurrent use
// We can use a lock to allow concurrent access of map
func mapConcurrent() {
	fmt.Println("\nMap Concurrent")
	set("a", "Some Data")
	result := get("a")
	fmt.Println(result)
}
