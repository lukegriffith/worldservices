package world

import "errors"

type worldService struct {
	Worlds       map[string]*World
	WorldsStatus map[string]Status
}

var (
	Worlds worldService
)

func (w worldService) GetWorldStatuses() map[string]string {
	statuses := make(map[string]string)
	for k, v := range w.WorldsStatus {
		statuses[k] = v.String()
	}
	return statuses
}

func NewWorldService() {
	Worlds = worldService{make(map[string]*World), make(map[string]Status)}
}

func GetWorldBoard(name string, cycle int) (GridHistory, error) {
	if world, ok := Worlds.Worlds[name]; ok {
		grid := world.GetCycle(cycle)
		return grid, nil
	}
	return GridHistory{}, errors.New("World not found")

}

func GetWorld(name string) World {
	return *Worlds.Worlds[name]
}

func RegisterWorld(name string) error {
	if _, ok := Worlds.WorldsStatus[name]; !ok {
		Worlds.WorldsStatus[name] = Pending
		return nil
	} else {
		return errors.New("World is already registered")
	}
}

func SetWorld(name string, w World) {
	if _, ok := Worlds.WorldsStatus[name]; ok {
		Worlds.Worlds[name] = &w
		Worlds.WorldsStatus[name] = Ready
	}
}
