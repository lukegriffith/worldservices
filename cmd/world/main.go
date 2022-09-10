package main

import (
	"fmt"

	"github.com/lukegriffith/worldservices/internal/server"
)

func main() {
	fmt.Println("Navigate to http://localhost:8080")
	server.SetupServer("8080", "./frontend")

}
