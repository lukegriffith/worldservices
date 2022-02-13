package main

import (
	"github.com/lukegriffith/worldservices"
)

func main() {
	worldservices.SetupServer("8080", "./frontend")
}
