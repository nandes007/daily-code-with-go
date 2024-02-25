package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/nandes007/daily-code-with-go/package-and-module-part1/math"
	"github.com/nandes007/daily-code-with-go/package-and-module-part1/math/advanced" //Nested package
	math2 "github.com/nandes007/daily-code-with-go/package-and-module-part1/math2"
)

func init() {
	fmt.Println("In main init")
}

func main() {
	uuidWithHypen, _ := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHypen.String(), "-", "", -1)
	fmt.Println(uuid)

	fmt.Println(add(2, 1))
	fmt.Println(subtract(2, 1))

	fmt.Println(math.Add(2, 1))
	fmt.Println(math.Subtract(2, 1))
	fmt.Println(math.Mul(2, 1))
	fmt.Println(advanced.Square(2))
	fmt.Println(math2.Subtract(2, 1))
}

func add(a, b int) int {
	return a + b
}
