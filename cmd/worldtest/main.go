package main

import (
	"github.com/lukegriffith/worldservices"
)

func main() {
	w := worldservices.NewWorld(100, 30)
	worldservices.WorldSingleton = &w
	worldservices.SetupServer("8080", "./frontend")
}
