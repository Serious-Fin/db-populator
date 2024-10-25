package main

import (
	"fmt"
	"log"
	parser "serious-fin/db-populator/ddl-parser"
)

func main() {
	fmt.Println("Hello, world!")
	_, err := parser.Parse("ola")
	if err != nil {
		log.Fatalf("An error occurred. Error: %v", err)
	}
}
