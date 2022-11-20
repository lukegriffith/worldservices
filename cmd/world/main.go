package main

import (
	"context"

	"github.com/lukegriffith/worldservices/internal/terminal"
	"github.com/lukegriffith/worldservices/internal/world"
)

func main() {
	var ctx = context.Background()

	world.NewWorldService()
	go terminal.SetupAndRun(ctx)
}
