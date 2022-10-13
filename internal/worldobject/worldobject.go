package worldobject

import (
	"math"
	"math/rand"
	"time"

	"github.com/patrikeh/go-deep"
)

type WorldObject interface {
	Fitness() float64
	GetCoordsXY() (int, int)
	Process(interface{}, float64)
	SetDebug()
	GetBrain() *deep.Neural
	SetBrain(d *deep.Neural)
}

type Stats struct {
	Health   int
	Strength int
	// Focus is used as vision distance.
	Focus int
	Speed int
	Age   int
}

func RandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func NewRandomStats() Stats {

	return Stats{
		RandomNumber(60, 100),
		RandomNumber(1, 10),
		RandomNumber(10, 25),
		RandomNumber(1, 4),
		RandomNumber(0, 80),
	}
}

func UpdateDistanceNeuron(distance float64, neuron *float64) {
	absDist := math.Abs(distance)
	if absDist == 1 {
		*neuron = 1.0
	} else {
		n1 := 1 - absDist/10
		if n1 > *neuron {
			*neuron = n1
		}
	}
}

func MinMax(array []float64) (int, int) {
	var max float64 = array[0]
	var maxIndex int
	var min float64 = array[0]
	var minIndex int
	for i, value := range array {
		if max < value {
			max = value
			maxIndex = i
		}
		if min > value {
			min = value
			minIndex = i
		}
	}
	return minIndex, maxIndex
}
