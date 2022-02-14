package main

import (
	"fmt"

	"github.com/lukegriffith/worldservices"
	"github.com/patrikeh/go-deep"
)

func main() {
	// Example of using crossover to create a new creature.
	c1 := worldservices.NewNormalCreature(1, 1)
	c2 := worldservices.NewNormalCreature(1, 2)

	b1 := c1.Net
	b2 := c2.Net
	// Could also do every other layer interlaced.
	crossoverPoint := len(b1.Layers) / 2
	l1 := b1.Layers[:len(b1.Layers)/crossoverPoint]
	l2 := b2.Layers[len(b2.Layers)/crossoverPoint:]

	l3 := append(l1, l2...)

	c3gen2 := worldservices.NormalCreature{
		S: worldservices.NewRandomStats(),
		X: 2,
		Y: 2,
		Net: &deep.Neural{
			Layers: l3,
			Biases: b2.Biases,
			Config: b2.Config,
		},
		LastControlSequence: nil,
		LastInputNeurons:    nil,
		Debug:               false,
	}
	objects := []worldservices.WorldObject{c1, c2}
	fmt.Println(c3gen2.Sense(objects, 1.0, 1.0))

	fmt.Println(len(c1.Net.Biases))
	fmt.Println(len(c1.Net.Layers))

	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c[:len(c)/2])
	fmt.Println(c[len(c)/2:])
}
