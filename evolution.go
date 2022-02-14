package worldservices

import "github.com/patrikeh/go-deep"

// crossover point
// select a random crossover point (range 0, len(Net.Layers))
// X & Y is new creature location
func CrossoverCreatures(c1 *NormalCreature, c2 *NormalCreature, X int, Y int) NormalCreature {

	n1 := c1.Net
	n2 := c2.Net
	// Could also do every other layer interlaced.

	crossoverPoint := randomNumber(0, len(n1.Layers))
	l1 := n1.Layers[:crossoverPoint]
	l2 := n2.Layers[crossoverPoint:]
	b1 := n1.Biases[:crossoverPoint]
	b2 := n1.Biases[crossoverPoint:]

	l3 := append(l1, l2...)
	b3 := append(b1, b2...)

	c3gen2 := NormalCreature{
		S: NewRandomStats(),
		X: X,
		Y: Y,
		Net: &deep.Neural{
			Layers: l3,
			Biases: b3,
			Config: n2.Config,
		},
		LastControlSequence: nil,
		LastInputNeurons:    nil,
		Debug:               false,
	}
	return c3gen2
}
