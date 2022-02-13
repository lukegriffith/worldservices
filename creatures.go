package worldservices

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/patrikeh/go-deep"
)

type WorldObject interface {
	Fitness() float64
	GetCoorsXY() (int, int)
	Process(Grid)
}

type BasicCreature struct {
	S                   Stats
	X, Y                int
	net                 *deep.Neural
	LastControlSequence []float64
}

type Stats struct {
	Health   int
	Strength int
	// Focus is used as vision distance.
	Focus int
	Speed int
	Age   int
}

func randomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func NewRandomStats() Stats {
	return Stats{
		randomNumber(60, 100),
		randomNumber(1, 10),
		randomNumber(1, 10),
		randomNumber(1, 10),
		randomNumber(0, 80),
	}
}

func (b *BasicCreature) Fitness() float64 {
	var fitness float64
	fitness = float64((b.S.Health * b.S.Speed) - b.S.Age)
	return fitness
}

func (b *BasicCreature) GetCoorsXY() (int, int) {
	return b.X, b.Y
}

func updateDistanceNeuron(distance float64, neuron *float64) {
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

// Some how this determines sensory data for distance. It wasn't thought
// much about.
func (b *BasicCreature) Sense(objects []WorldObject) []float64 {
	bX, bY := b.GetCoorsXY()
	var xPlusNeuron, xMinusNeuron, yPlusNeuron, yMinusNeuron float64
	for _, obj := range objects {
		objX, objY := obj.GetCoorsXY()
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

func minMax(array []float64) (int, int) {
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

func (b *BasicCreature) Process(g Grid) {
	// board input.
	// build inputs from grid and creature
	sensedObjects := g.GetObjectSenseData(b.X, b.Y, b.S.Focus)
	neuralInput := b.Sense(sensedObjects)
	fmt.Println("neural input", neuralInput)
	controlSequence := b.net.Predict(neuralInput)
	fmt.Println("control sequence", controlSequence)
	b.LastControlSequence = controlSequence

	_, largestIndex := minMax(controlSequence)

	if largestIndex == 0 {
		fmt.Printf("MOVED new%d old:%d\n", b.X, b.X+1)
		fmt.Println(b)
		b.X = b.X + 1
		fmt.Println(b)

	}
	if largestIndex == 1 {
		fmt.Printf("MOVED new%d old:%d\n", b.X, b.X+1)
		fmt.Println(b)
		b.X = b.X + 1
		fmt.Println(b)
	}
	if largestIndex == 2 {
		fmt.Printf("MOVED new%d old:%d\n", b.Y, b.Y+1)
		fmt.Println(b)
		b.Y = b.Y + 1
		fmt.Println(b)
	}
	if largestIndex == 3 {
		fmt.Printf("MOVED new%d old:%d\n", b.Y, b.Y+1)
		fmt.Println(b)
		b.Y = b.Y - 1
		fmt.Println(b)
	}

	// Output is control sequence.
}

func NewBasicCreature(x int, y int) *BasicCreature {
	// Initial chromosones
	// Cross over if bread
	n := createNetwork(4, []int{2, 2, 4})
	trainNetwork(n, BasicTraining)
	// Train network based on chromosones
	return &BasicCreature{NewRandomStats(), x, y, n, nil}
}
