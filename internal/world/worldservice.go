package world

import "errors"

type worldService map[string]*World

var (
	Worlds worldService
)

func NewWorldService() {
	Worlds = make(worldService)
}

func GetWorldBoard(name string, cycle int) (GridHistory, error) {
	if world, ok := Worlds[name]; ok {
		grid := world.GetCycle(cycle)
		return grid, nil
	}
	return GridHistory{}, errors.New("World not found")

}

func GetWorld(name string) World {
	return *Worlds[name]
}

func RegisterWorld(name string, w World) {
	Worlds[name] = &w
}
