package main

import (
	"fmt"
	"os"
)

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error One happended"
}

// Use case wrapped the errors
type notPositive struct {
	num int
}

func (e notPositive) Error() string {
	return fmt.Sprintf("checkPositive: Given number %d is not a positive number", e.num)
}

type notEven struct {
	num int
}

func (e notEven) Error() string {
	return fmt.Sprintf("checkEven: Given number %d is not an even number", e.num)
}

func checkPositive(num int) error {
	if num < 0 {
		return notPositive{num: num}
	}
	return nil
}

func checkEven(num int) error {
	if num%2 == 1 {
		return notEven{num: num}
	}
	return nil
}

func checkPositiveAndEven(num int) error {
	if num > 100 {
		return fmt.Errorf("checkPositiveAndEven: Number %d is greather than 100", num)
	}

	err := checkPositive(num)
	if err != nil {
		return fmt.Errorf("checkPositiveAndEven: %w", err)
	}

	err = checkEven(num)
	if err != nil {
		return fmt.Errorf("checkPositiveAndEven: %w", err)
	}

	return nil
}

func do() error {
	return fmt.Errorf("E2: %w", errorOne{})
}

func openFile(fileName string) error {
	_, err := os.Open(fileName)
	if err != nil {
		return err
	}
	return nil
}

func openFileWrappedError(fileName string) error {
	_, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Error opening: %w", err)
	}
	return nil
}
