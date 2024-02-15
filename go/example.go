package main

import (
	"fmt"
	"log"

	"golang.org/x/text/encoding"
)

func main() {
	nop := encoding.Nop
	encoder := nop.NewEncoder()
	result, err := encoder.String("Hello")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Print(result)
	fmt.Print("Hello, World!\n")
}
