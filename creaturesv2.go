package worldservices

import (
	"math"

	"github.com/patrikeh/go-deep"
)

type NormalCreature struct {
	S                   Stats
	X, Y                int
	net                 *deep.Neural
	LastControlSequence []float64
	LastInputNeurons    []float64
	Debug               bool
}

func (b *NormalCreature) Fitness() float64 {
	fitness := float64((b.S.Health * b.S.Speed) - b.S.Age)
	return fitness
}

func (b *NormalCreature) GetCoordsXY() (int, int) {
	return b.X, b.Y
}

func updateDistanceNeuronV2(distance float64, neuron *float64) {
	absDist := math.Abs(distance)
	*neuron = absDist + *neuron
}

// Some how this determines sensory data for distance. It wasn't thought
// much about.
func (b *NormalCreature) Sense(objects []WorldObject, oscilator float64, age float64) []float64 {

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
	return []float64{xPlusNeuron, xMinusNeuron, yPlusNeuron, yMinusNeuron, oscilator, age}
}

func (b *NormalCreature) Process(g Grid, oscilator float64) {
	// board input.
	// build inputs from grid and creature
	sensedObjects := g.GetObjectSenseData(b.X, b.Y, b.S.Focus)
	neuralInput := b.Sense(sensedObjects, oscilator, float64(b.S.Age))
	network := b.GetBrain()
	controlSequence := network.Predict(neuralInput)
	b.LastControlSequence = controlSequence
	b.LastInputNeurons = neuralInput
	b.S.Age = b.S.Age + 1

	_, largestIndex := minMax(controlSequence)

	value := controlSequence[largestIndex]
	// Added ability for a creature to stay still if no neuron fires above .50
	stride := int(float64(b.S.Speed) * value)
	if largestIndex == 0 {
		newX := b.X + stride
		_, err := g.GetObjectAtCoords(newX, b.Y)
		if err != nil && newX < g.Size {
			b.X = newX
		}
	}
	if largestIndex == 1 {
		newX := b.X - stride
		_, err := g.GetObjectAtCoords(newX, b.Y)

		if err != nil && newX > 0 {
			b.X = newX
		}
	}
	if largestIndex == 2 {
		newY := b.Y + stride
		_, err := g.GetObjectAtCoords(b.X, newY)
		if err != nil && newY < g.Size {
			b.Y = newY
		}
	}
	if largestIndex == 3 {
		newY := b.Y - stride
		_, err := g.GetObjectAtCoords(newY, b.Y-1)
		if err != nil && newY > 0 {
			b.Y = newY
		}
	}
	g.UpdateLocationsCoords()

}

func (b *NormalCreature) SetDebug() {
	b.Debug = !b.Debug
}

func (b *NormalCreature) GetBrain() *deep.Neural {
	return b.net
}

func (b *NormalCreature) SetBrain(d *deep.Neural) {
	b.net = d
}

func NewNormalCreature(x int, y int) *NormalCreature {
	// Initial chromosones
	// Cross over if bread
	n := createNetwork(6, []int{2, 2, 4})
	//trainNetwork(n, BasicTrainingWOscilationAndAge)
	// Train network based on chromosones
	return &NormalCreature{NewRandomStats(), x, y, n, nil, nil, false}
}
