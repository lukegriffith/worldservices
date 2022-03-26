package worldservices

type worldService map[string]*World

var (
	Worlds worldService
)

func NewWorldService() {
	Worlds = make(worldService)
}

func GetWorldBoard(name string, cycle int) GridHistory {
	world := Worlds[name]
	grid := world.GetCycle(cycle)
	return grid
}

func GetWorld(name string) World {
	return *Worlds[name]
}

func RegisterWorld(name string, w World) {
	Worlds[name] = &w
}
