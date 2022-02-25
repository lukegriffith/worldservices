package worldservices

import (
	"time"
)

type World struct {
	Grid    Grid
	cycleNo int
	History WorldHistory
}

func generateSafeLocation(locations map[string]WorldObject, size int) (int, int) {
	x := randomNumber(0, size)
	time.Sleep(3 * time.Nanosecond)
	y := randomNumber(0, size)
	_, exists := locations[formatCoords(x, y)]
	// TODO: validate pop size is lower than size ** size
	// this is a very costly way of doing it.
	for exists {
		x := randomNumber(0, size)
		time.Sleep(3 * time.Nanosecond)
		y := randomNumber(0, size)
		_, exists = locations[formatCoords(x, y)]
	}
	return x, y
}

func NewWorld(size int, pop int) World {
	creatures := []WorldObject{}
	locations := map[string]WorldObject{}
	for i := 0; i <= pop; i++ {
		x, y := generateSafeLocation(locations, size)
		creature := NewNormalCreature(x, y)
		locations[formatCoords(x, y)] = creature
		creatures = append(creatures, creature)
	}
	return World{Grid{creatures, locations, size}, 0}
}

func (w *World) Oscilator() float64 {
	return float64(w.cycleNo%10) / 10
}

func (w *World) Cycle() {
	w.History.Push(w.Grid)
	objects := w.Grid.GetOrderedObjectListByFitness()
	for _, o := range objects {
		o.Process(w.Grid, w.Oscilator())
	}
	w.Grid.UpdateLocationsCoords()
	w.cycleNo = w.cycleNo + 1
}

// NewWorldFromDebug
// creates a new world from the creatures in debug mode
// using the crossover function.
func NewWorldFromDebug(world string) *World {
	w := GetWorld(world)
	nextGeneration := []WorldObject{}
	debuggedObjects := []WorldObject{}
	objects := w.Grid.GetOrderedObjectListByFitness()

	for _, obj := range objects {
		c := obj.(*NormalCreature)
		if c.Debug {
			debuggedObjects = append(debuggedObjects, c)
		}
	}
	var locations map[string]WorldObject
	for len(debuggedObjects) > 1 {
		c1 := debuggedObjects[0].(*NormalCreature)
		debuggedObjects = debuggedObjects[1:]
		c2 := debuggedObjects[0].(*NormalCreature)
		debuggedObjects = debuggedObjects[1:]
		X, Y := generateSafeLocation(locations, w.Grid.Size)
		c3gen2, _ := CrossoverCreatures(c1, c2, X, Y)
		nextGeneration = append(nextGeneration, &c3gen2)
	}
	return &World{
		Grid: Grid{
			objects:   nextGeneration,
			locations: nil,
			Size:      w.Grid.Size,
		},
		cycleNo: 0,
	}
}

func (w *World) Run(simLength int) {
	for i := 1; i < simLength; i++ {
		w.Cycle()
	}
}
