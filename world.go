package worldservices

import (
	"errors"
	"time"
)

var (
	WorldSingleton *World
)

func GetWorldSingleton() (*World, error) {
	if WorldSingleton != nil {
		return WorldSingleton, nil
	}
	return nil, errors.New("Singleton not set.")
}

type World struct {
	Grid Grid
}

func NewWorld(size int, pop int) World {
	creatures := []WorldObject{}
	locations := map[string]WorldObject{}
	for i := 0; i <= pop; i++ {
		x := randomNumber(0, size)
		time.Sleep(3 * time.Nanosecond)
		y := randomNumber(0, size)
		_, exists := locations[formatCoords(x, y)]
		for exists {
			x := randomNumber(0, size)
			time.Sleep(3 * time.Nanosecond)
			y := randomNumber(0, size)
			_, exists = locations[formatCoords(x, y)]
		}

		creature := NewNormalCreature(x, y)

		locations[formatCoords(x, y)] = creature
		creatures = append(creatures, creature)
	}
	return World{Grid{creatures, locations, size}}
}

func (w *World) Cycle() {
	objects := w.Grid.GetOrderedObjectListByFitness()
	for _, o := range objects {
		o.Process(w.Grid)
	}
	w.Grid.UpdateLocationsCoords()
}
