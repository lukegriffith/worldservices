package worldservices

import (
	"math"

	"github.com/patrikeh/go-deep"
)

func updateDistanceNeuronV2(distance float64, neuron *float64) {
	absDist := math.Abs(distance)

	if distance > *neuron {
		*neuron = absDist
	}
}

type NormalCreature struct {
	S                   Stats
	X, Y                int
	net                 *deep.Neural
	LastControlSequence []float64
	LastInputNeurons    []float64
	Debug               bool
}

func (b *NormalCreature) Fitness() float64 {
	var fitness float64
	fitness = float64((b.S.Health * b.S.Speed) - b.S.Age)
	return fitness
}

func (b *NormalCreature) GetCoordsXY() (int, int) {
	return b.X, b.Y
}

// Some how this determines sensory data for distance. It wasn't thought
// much about.
func (b *NormalCreature) Sense(objects []WorldObject, oscilator float64) []float64 {
	bX, bY := b.GetCoordsXY()
	var xPlusNeuron, xMinusNeuron, yPlusNeuron, yMinusNeuron float64
	for _, obj := range objects {
		objX, objY := obj.GetCoordsXY()
		xdistance := float64(objX - bX)
		ydistance := float64(objY - bY)

		if xdistance > 0 {
			updateDistanceNeuronV2(xdistance, &xPlusNeuron)
		} else {
			updateDistanceNeuronV2(xdistance, &xMinusNeuron)
		}

		if ydistance > 0 {
			updateDistanceNeuronV2(ydistance, &yPlusNeuron)
		} else {
			updateDistanceNeuronV2(ydistance, &yMinusNeuron)
		}
	}
	return []float64{xPlusNeuron, xMinusNeuron, yPlusNeuron, yMinusNeuron, oscilator}
}

func (b *NormalCreature) Process(g Grid, oscilator float64) {
	// board input.
	// build inputs from grid and creature
	sensedObjects := g.GetObjectSenseData(b.X, b.Y, b.S.Focus)
	neuralInput := b.Sense(sensedObjects, oscilator)
	controlSequence := b.net.Predict(neuralInput)
	b.LastControlSequence = controlSequence
	b.LastInputNeurons = neuralInput
	b.S.Age = b.S.Age + 1

	_, largestIndex := minMax(controlSequence)

	value := controlSequence[largestIndex]
	// Added ability for a creature to stay still if no neuron fires above .50
	if value > 0.5 {
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
}

func (b *NormalCreature) SetDebug() {
	b.Debug = !b.Debug
}

func NewNormalCreature(x int, y int) *NormalCreature {
	// Initial chromosones
	// Cross over if bread
	n := createNetwork(5, []int{2, 2, 4})
	trainNetwork(n, BasicTrainingWOscilation)
	// Train network based on chromosones
	return &NormalCreature{NewRandomStats(), x, y, n, nil, nil, false}
}
