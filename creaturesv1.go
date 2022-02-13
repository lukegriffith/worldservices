package worldservices

import (
	"github.com/patrikeh/go-deep"
)

type BasicCreature struct {
	S                   Stats
	X, Y                int
	net                 *deep.Neural
	LastControlSequence []float64
}

func (b *BasicCreature) Fitness() float64 {
	var fitness float64
	fitness = float64((b.S.Health * b.S.Speed) - b.S.Age)
	return fitness
}

func (b *BasicCreature) GetCoordsXY() (int, int) {
	return b.X, b.Y
}

// Some how this determines sensory data for distance. It wasn't thought
// much about.
func (b *BasicCreature) Sense(objects []WorldObject) []float64 {
	bX, bY := b.GetCoordsXY()
	var xPlusNeuron, xMinusNeuron, yPlusNeuron, yMinusNeuron float64
	for _, obj := range objects {
		objX, objY := obj.GetCoordsXY()
		xdistance := float64(objX - bX)
		ydistance := float64(objY - bY)

		if xdistance > 0 {
			updateDistanceNeuron(xdistance, &xPlusNeuron)
		} else {
			updateDistanceNeuron(xdistance, &xMinusNeuron)
		}

		if ydistance > 0 {
			updateDistanceNeuron(ydistance, &yPlusNeuron)
		} else {
			updateDistanceNeuron(ydistance, &yMinusNeuron)
		}

	}
	return []float64{xPlusNeuron, xMinusNeuron, yPlusNeuron, yMinusNeuron}
}

func (b *BasicCreature) Process(g Grid) {
	// board input.
	// build inputs from grid and creature
	sensedObjects := g.GetObjectSenseData(b.X, b.Y, b.S.Focus)
	neuralInput := b.Sense(sensedObjects)
	controlSequence := b.net.Predict(neuralInput)
	b.LastControlSequence = controlSequence

	_, largestIndex := minMax(controlSequence)

	if largestIndex == 0 {
		_, err := g.GetObjectAtCoords(b.X+1, b.Y)
		if err != nil && b.X+1 < g.Size {
			b.X = b.X + 1
		}
	}
	if largestIndex == 1 {
		_, err := g.GetObjectAtCoords(b.X-1, b.Y)

		if err != nil && b.X-1 > 0 {
			b.X = b.X - 1
		}
	}
	if largestIndex == 2 {
		_, err := g.GetObjectAtCoords(b.X, b.Y+1)
		if err != nil && b.Y+1 < g.Size {
			b.Y = b.Y + 1
		}
	}
	if largestIndex == 3 {
		_, err := g.GetObjectAtCoords(b.X, b.Y-1)
		if err != nil && b.Y-1 > 0 {
			b.Y = b.Y - 1
		}
	}
	g.UpdateLocationsCoords()
}

func NewBasicCreature(x int, y int) *BasicCreature {
	// Initial chromosones
	// Cross over if bread
	n := createNetwork(4, []int{2, 2, 4})
	trainNetwork(n, BasicTraining)
	// Train network based on chromosones
	return &BasicCreature{NewRandomStats(), x, y, n, nil}
}
