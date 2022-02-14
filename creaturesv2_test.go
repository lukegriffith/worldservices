package worldservices_test

import (
	"testing"

	"github.com/lukegriffith/worldservices"
)

func TestCreatureSense(t *testing.T) {
	objects := []worldservices.WorldObject{
		worldservices.NewNormalCreature(3, 3),
		worldservices.NewNormalCreature(3, 2),
		worldservices.NewNormalCreature(12, 20),
	}
	g := SetupTestGrid(objects)
	obj, err := g.GetObjectAtCoords(3, 3)
	if err != nil {
		t.FailNow()
	}
	nc := obj.(*worldservices.NormalCreature)
	sensed := g.GetObjectSenseData(3, 3, 3)
	neuronOutput := nc.Sense(sensed, 0.9, 1.0)

	t.Log(neuronOutput)
}
