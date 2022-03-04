package worldservices

var (
	Worlds map[string]World
)

func GetWorldBoard(name string, cycle int) Grid {
	grid := Worlds[name].GetCycle(cycle)
	return grid
}

func GetWorld(name string) World {
	return Worlds[name]
}

func RegisterWorld(name string, w World) {
	Worlds[name] = w
}
