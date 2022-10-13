package ml

import (
	"github.com/patrikeh/go-deep"

	"github.com/lukegriffith/worldservices/internal/worldobject"
)

// TODO: Need to implement crossover.
// crossover point
// select a random crossover point (range 0, len(Net.Layers))
// X & Y is new creature location
// returned int is crossover point.
func CrossoverCreatures(c1 worldobject.WorldObject, c2 worldobject.WorldObject, c3 worldobject.WorldObject) (worldobject.WorldObject, int) {

	n1 := c1.GetBrain()
	n2 := c2.GetBrain()
	// Could also do every other layer interlaced.

	crossoverPoint := worldobject.RandomNumber(0, len(n1.Layers))
	l1 := n1.Layers[:crossoverPoint]
	l2 := n2.Layers[crossoverPoint:]
	b1 := n1.Biases[:crossoverPoint]
	b2 := n1.Biases[crossoverPoint:]

	l3 := append(l1, l2...)
	b3 := append(b1, b2...)
	brain := &deep.Neural{
		Layers: l3,
		Biases: b3,
		Config: n2.Config,
	}
	c3.SetBrain(brain)
	return c3, crossoverPoint
}
