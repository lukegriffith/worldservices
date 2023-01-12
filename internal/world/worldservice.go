package world

var (
	worldIndex int32 = 0
)

type WorldsContainer struct {
	Worlds      []WorldsObject
	ActiveWorld int32
}

type WorldsObject struct {
	Id     int32
	Name   string
	World  *World
	Status Status
}

func (w *WorldsContainer) NewWorld(Name string) int32 {
	worldObj := WorldsObject{
		Id:     worldIndex,
		Name:   Name,
		World:  nil,
		Status: Pending,
	}
	w.Worlds = append(w.Worlds, worldObj)
	worldIndex++
	return worldIndex
}

/*
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
*/
