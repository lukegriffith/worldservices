package main

import (
	"fmt"

	"github.com/lukegriffith/worldservices"
)

func main() {
	fmt.Println("Navigate to http://localhost:8080")
	worldservices.SetupServer("8080", "./frontend")

}
