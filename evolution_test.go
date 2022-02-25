package worldservices_test

import (
	"testing"

	"github.com/lukegriffith/worldservices"
)

func TestCrossoverCreature(t *testing.T) {

	c1 := worldservices.NewNormalCreature(3, 3)
	c2 := worldservices.NewNormalCreature(3, 2)

	_, _ = worldservices.CrossoverCreatures(c1, c2, 0, 0)
	/*
		n := c.GetBrain()
		l1 := n.Layers
	*/

}
