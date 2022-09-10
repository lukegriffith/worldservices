package world

import (
	"time"

	"github.com/lukegriffith/worldservices/internal/creatures/v2"
	critters "github.com/lukegriffith/worldservices/internal/creatures/v2"
	"github.com/lukegriffith/worldservices/internal/grid"
	"github.com/lukegriffith/worldservices/internal/ml"
	"github.com/lukegriffith/worldservices/internal/worldobject"
)

type World struct {
	Grid    grid.Grid
	cycleNo int
	history *WorldHistory
}

func generateSafeLocation(locations map[string]worldobject.WorldObject, size int) (int, int) {
	x := worldobject.RandomNumber(0, size)
	time.Sleep(3 * time.Nanosecond)
	y := worldobject.RandomNumber(0, size)
	_, exists := locations[grid.FormatCoords(x, y)]
	// TODO: validate pop size is lower than size ** size
	// this is a very costly way of doing it.
	for exists {
		x := worldobject.RandomNumber(0, size)
		time.Sleep(3 * time.Nanosecond)
		y := worldobject.RandomNumber(0, size)
		_, exists = locations[grid.FormatCoords(x, y)]
	}
	return x, y
}

func NewWorld(size int, pop int) World {
	creatures := []worldobject.WorldObject{}
	locations := map[string]worldobject.WorldObject{}
	for i := 0; i <= pop; i++ {
		x, y := generateSafeLocation(locations, size)
		creature := critters.NewNormalCreature(x, y)
		locations[grid.FormatCoords(x, y)] = creature
		creatures = append(creatures, creature)
	}
	return World{grid.NewGrid(creatures, locations, size), 0, &WorldHistory{}}
}

func (w *World) Oscilator() float64 {
	return float64(w.cycleNo%10) / 10
}

func (w *World) Cycle() {
	w.history.Push(w.Grid)
	objects := w.Grid.GetOrderedObjectListByFitness()
	for _, o := range objects {
		o.Process(w.Grid, w.Oscilator())
	}
	w.Grid.UpdateLocationsCoords()
	w.cycleNo = w.cycleNo + 1
	w.Grid.SetCycle(w.cycleNo)
}

func (w *World) GetCycle(cycle int) GridHistory {
	g := w.history.Get(cycle)
	return g
}

func (w *World) Run(simLength int) {
	for i := 1; i < simLength; i++ {
		w.Cycle()
	}
}

// NewWorldFromDebug
// creates a new world from the creatures in debug mode
// using the crossover function.
func NewWorldFromDebug(objects []worldobject.WorldObject, newWorldSize int) World {

	nextGeneration := []worldobject.WorldObject{}
	debuggedObjects := []worldobject.WorldObject{}

	for _, obj := range objects {
		c := obj.(*creatures.NormalCreature)
		if c.Debug {
			debuggedObjects = append(debuggedObjects, c)
		}
	}
	var locations map[string]worldobject.WorldObject
	for len(debuggedObjects) > 1 {
		c1 := debuggedObjects[0].(*creatures.NormalCreature)
		debuggedObjects = debuggedObjects[1:]
		c2 := debuggedObjects[0].(*creatures.NormalCreature)
		debuggedObjects = debuggedObjects[1:]
		X, Y := generateSafeLocation(locations, newWorldSize)
		c3 := creatures.NewNormalCreature(X, Y)
		c3gen2, _ := ml.CrossoverCreatures(c1, c2, c3)
		c3gen2crit := c3gen2.(creatures.NormalCreature)
		nextGeneration = append(nextGeneration, &c3gen2crit)
	}
	return World{
		Grid:    grid.NewGrid(nextGeneration, nil, newWorldSize),
		cycleNo: 0,
	}
}

func WorldFromDebugOfWorlds(g1 grid.Grid, g2 grid.Grid) World {
	objects := append(g1.GetObjects(), g2.GetObjects()...)
	return NewWorldFromDebug(objects, g1.Size)
}
