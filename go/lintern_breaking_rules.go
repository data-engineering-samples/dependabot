package main

import (
	"errors"
	"fmt"
)

// MyStruct represents a sample struct.
type MyStruct struct {
	// golint error: should be named myField
	Myfield int
}

// MyFunction is a sample exported function.
func MyFunction() {
	// golint error: error strings should not be capitalized or end with punctuation or a newline
	errors.New("This is an error.")
}

// MyMethod is a sample method.
// golint error: receiver name should not be an underscore, omit the name if it is unused
func (_ *MyStruct) MyMethod() {
	// golint error: error strings should not be capitalized or end with punctuation or a newline
	errors.New("This is another error.")
}

// main is the entry point for the program.
func main() {
	// golint error: unused import
	fmt.Println("Hello, world!")
}
