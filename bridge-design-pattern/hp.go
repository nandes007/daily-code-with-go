package main

import "fmt"

type hp struct {
}

func (h *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}
