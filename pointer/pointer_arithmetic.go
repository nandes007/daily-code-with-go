package main

func main() {
	//Pointer arithmetic is not possible in golang unlike C language. It raises compilation error.
	a := 1
	b := &a
	b = b + 1 //This line will raise error.
}
